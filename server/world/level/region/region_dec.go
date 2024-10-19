package region

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"sync"
	"unsafe"

	"github.com/4kills/go-zlib"
	"github.com/aimjel/minecraft/nbt"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/compress"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
)

type File struct {
	rx, rz int32

	generator chunk.Generator

	reader    io.ReaderAt
	locations []byte

	chunks map[uint64]*chunk.Chunk
	chu_mu sync.Mutex
}

func chunkLocation(l int32) (offset, size int32) {
	offset = ((l >> 8) & 0xFFFFFF)
	size = l & 0xFF

	return offset * 4096, size * 4096
}

func (r *File) generateChunkAt(x, z int32, tgt *chunk.Chunk, generator chunk.Generator) {
	c := generator.NewChunk(x, z)

	*tgt = c
}

func (r *File) LoadedChunks() int32 {
	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	return int32(len(r.chunks))
}

func (r *File) GetChunk(x, z int32) (*chunk.Chunk, error) {
	var chunkBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	chunkBuffer.Reset()
	defer buffers.Buffers.Put(chunkBuffer)

	return r.GetChunkBuf(x, z, chunkBuffer)
}

func (r *File) GetChunkBuf(x, z int32, chunkBuffer *bytes.Buffer) (*chunk.Chunk, error) {
	hash := ChunkHash(x, z)

	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	if c, ok := r.chunks[hash]; ok {
		return c, nil
	}

	locationIndex := 4 * ((x & 31) + (z&31)*32)
	if int(locationIndex) >= len(r.locations) {
		if r.generator != nil {
			r.chunks[hash] = new(chunk.Chunk)
			r.generateChunkAt(x, z, r.chunks[hash], r.generator)
			return r.chunks[hash], nil
		}
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	l := r.locations[locationIndex : locationIndex+4]
	loc := int32(l[0])<<24 | int32(l[1])<<16 | int32(l[2])<<8 | int32(l[3])

	offset, size := chunkLocation(loc)
	if offset == 0 && size == 0 {
		if r.generator != nil {
			r.chunks[hash] = new(chunk.Chunk)
			r.generateChunkAt(x, z, r.chunks[hash], r.generator)
			return r.chunks[hash], nil
		}
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	var chunkHeader [5]byte

	_, err := r.reader.ReadAt(chunkHeader[:], int64(offset))
	if err != nil {
		return nil, err
	}

	length := (int32(chunkHeader[0])<<24 | int32(chunkHeader[1])<<16 | int32(chunkHeader[2])<<8 | int32(chunkHeader[3])) - 1
	compression := chunkHeader[4]

	if length == 0 {
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	var reader = io.NewSectionReader(r.reader, int64(offset)+5, int64(length))

	switch compression {
	case CompressionZlib:
		z := compress.RZlib.Get().(*zlib.Reader)
		z.Reset(reader, nil)

		chunkBuffer.ReadFrom(z)
		compress.RZlib.Put(z)
	case CompressionGzip:
		g := compress.RGzip.Get().(*gzip.Reader)
		g.Reset(reader)

		chunkBuffer.ReadFrom(g)
		compress.RGzip.Put(z)
	case CompressionNone:
		chunkBuffer.ReadFrom(reader)
	}

	var anvil = new(anvilChunk)

	if err := nbt.Unmarshal(chunkBuffer.Bytes(), anvil); err != nil {
		return nil, err
	}

	r.chunks[hash] = &chunk.Chunk{
		X:             anvil.XPos,
		Y:             anvil.YPos,
		Z:             anvil.ZPos,
		Heightmaps:    anvil.Heightmaps,
		BlockEntities: anvil.BlockEntities,
	}

	r.chunks[hash].Sections = make([]*section.Section, len(anvil.Sections))
	var emptySections int
	for i, sec := range anvil.Sections {
		var blocks = make([]section.Block, len(sec.BlockStates.Palette))

		if l := len(sec.BlockStates.Palette); l == 0 || (l == 1 && sec.BlockStates.Palette[0].Name == "minecraft:air") {
			emptySections++
		}
		for i, entry := range sec.BlockStates.Palette {
			b := section.GetBlock(entry.Name)
			if entry.Properties != nil {
				b = b.New(entry.Properties)
			}
			blocks[i] = b
		}

		r.chunks[hash].Sections[i] = section.New(sec.Y,
			blocks,
			sec.BlockStates.Data,
			sec.Biomes.Palette,
			sec.Biomes.Data,
			*(*[]byte)(unsafe.Pointer(&sec.SkyLight)),
			*(*[]byte)(unsafe.Pointer(&sec.BlockLight)),
		)
	}
	/*if emptySections == len(r.chunks[hash].Sections) && r.generateEmpty && r.generator != nil {
		r.chunks[hash] = new(chunk.Chunk)
		r.generateChunkAt(x, z, r.chunks[hash], r.generator)
		return r.chunks[hash], nil
	}*/

	return r.chunks[hash], err
}

func Empty(f *File, rx, rz int32, generator chunk.Generator) {
	*f = File{
		chunks: make(map[uint64]*chunk.Chunk),
		rx:     rx, rz: rz,
		generator: generator,
	}
}

func Decode(r io.ReaderAt, f *File, rx, rz int32, generator chunk.Generator) error {
	var locationTable = make([]byte, 4096)

	_, err := r.ReadAt(locationTable, 0)
	if err != nil {
		return err
	}

	*f = File{
		chunks: make(map[uint64]*chunk.Chunk),
		rx:     rx, rz: rz,
		generator: generator,
		locations: locationTable,
		reader:    r,
	}

	return nil
	//return f.decodeAll(locationTable, r)
}

func ChunkHash(x, z int32) uint64 {
	return uint64(uint32(z))<<32 | uint64(uint32(x))
}
