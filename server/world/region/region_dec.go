package region

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"sync"

	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/net/buffers"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
)

type Generator interface {
	NewChunk(x, z int32) chunk.Chunk
}

type File struct {
	reader io.ReaderAt

	locations []byte

	chunks map[uint64]*chunk.Chunk
	chu_mu sync.Mutex
}

func chunkLocation(l int32) (offset, size int32) {
	offset = ((l >> 8) & 0xFFFFFF)
	size = l & 0xFF

	return offset * 4096, size * 4096
}

func (r *File) generateChunkAt(x, z int32, tgt *chunk.Chunk, generator Generator) {
	c := generator.NewChunk(x, z)

	*tgt = c
}

func (r *File) LoadedChunks() int32 {
	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	return int32(len(r.chunks))
}

func (r *File) GetChunk(x, z int32, generateEmpty bool, generator Generator) (*chunk.Chunk, error) {
	hash := ChunkHash(x, z)

	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	if c, ok := r.chunks[hash]; ok {
		return c, nil
	}

	//r.chunks[hash] = new(chunk.Chunk)
	//r.generateChunkAt(x, z, r.chunks[hash], generator)
	//return r.chunks[hash], nil

	locationIndex := ((uint32(x) % 32) + (uint32(z)%32)*32) * 4
	if int(locationIndex) >= len(r.locations) {
		if generator != nil {
			r.chunks[hash] = new(chunk.Chunk)
			r.generateChunkAt(x, z, r.chunks[hash], generator)
			return r.chunks[hash], nil
		}
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	l := r.locations[locationIndex : locationIndex+4]
	loc := int32(l[0])<<24 | int32(l[1])<<16 | int32(l[2])<<8 | int32(l[3])

	offset, size := chunkLocation(loc)
	if offset == 0 && size == 0 {
		if generator != nil {
			r.chunks[hash] = new(chunk.Chunk)
			r.generateChunkAt(x, z, r.chunks[hash], generator)
			return r.chunks[hash], nil
		}
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	var chunkHeader = make([]byte, 5)

	_, err := r.reader.ReadAt(chunkHeader, int64(offset))
	if err != nil {
		return nil, err
	}

	length := int32(chunkHeader[0])<<24 | int32(chunkHeader[1])<<16 | int32(chunkHeader[2])<<8 | int32(chunkHeader[3])
	compression := chunkHeader[4]

	if length == 0 {
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	var chunkData = make([]byte, length-1)
	_, err = r.reader.ReadAt(chunkData, int64(offset)+5)
	if err != nil {
		return nil, err
	}

	var rd io.ReadCloser

	switch compression {
	case 1:
		rd, err = gzip.NewReader(bytes.NewReader(chunkData))
		if err != nil {
			return nil, err
		}
		defer rd.Close()
	case 2:
		rd, err = zlib.NewReader(bytes.NewReader(chunkData))
		if err != nil {
			return nil, err
		}
		defer rd.Close()
	}

	buf := buffers.Buffers.Get().(*bytes.Buffer)
	buf.Reset()
	buf.ReadFrom(rd)
	defer buffers.Buffers.Put(buf)

	var anvil anvilChunk

	_, err = nbt.NewDecoder(buf).Decode(&anvil)

	if err != nil {
		fmt.Println(err)
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

		r.chunks[hash].Sections[i] = section.New(sec.Y, blocks, sec.BlockStates.Data, sec.Biomes.Palette, sec.Biomes.Data, sec.SkyLight, sec.BlockLight)
	}

	if emptySections == len(r.chunks[hash].Sections) && generateEmpty && generator != nil {
		r.chunks[hash] = new(chunk.Chunk)
		r.generateChunkAt(x, z, r.chunks[hash], generator)
		return r.chunks[hash], nil
	}

	return r.chunks[hash], err

	/*chunk, ok := r.chunks[loc]
		if !ok {
			return chunk, fmt.Errorf("not found chunk")
		}
		return chunk, nil
	}
	*/
}

func Empty(f *File) {
	*f = File{
		chunks: make(map[uint64]*chunk.Chunk),
	}
}

func Decode(r io.ReaderAt, f *File) error {
	var locationTable = make([]byte, 4096)

	_, err := r.ReadAt(locationTable, 0)
	if err != nil {
		return err
	}

	*f = File{
		reader: r,

		locations: locationTable,
		chunks:    make(map[uint64]*chunk.Chunk),
	}

	/*var chunkBuffer = new(bytes.Buffer)

	for i := 0; i < 1024; i++ {
		loc := int32(locationTable[(i*4)+0])<<24 | int32(locationTable[(i*4)+1])<<16 | int32(locationTable[(i*4)+2])<<8 | int32(locationTable[(i*4)+3])

		offset, size := chunkLocation(loc)
		if offset == 0 && size == 0 {
			continue
		}
		var chunkHeader [5]byte
		if _, err := r.ReadAt(chunkHeader[:], int64(offset)); err != nil {
			return err
		}

		var length = binary.BigEndian.Uint32(chunkHeader[:4]) - 1
		var compressionScheme = chunkHeader[4]

		var chunkData = make([]byte, length-1)

		_, err = r.ReadAt(chunkData, int64(offset)+5)
		if err != nil {
			return err
		}

		var rd io.ReadCloser

		switch compressionScheme {
		case 1:
			rd, err = gzip.NewReader(bytes.NewReader(chunkData))
			if err != nil {
				return err
			}
			defer rd.Close()
		case 2:
			rd, err = zlib.NewReader(bytes.NewReader(chunkData))
			if err != nil {
				return err
			}
			defer rd.Close()
		}

		chunkBuffer.Reset()
		chunkBuffer.ReadFrom(rd)

		f.chunks[loc] = &chunk.Chunk{}

		_, err = nbt.NewDecoder(chunkBuffer).Decode(f.chunks[loc])
		if err != nil {
			return err
		}
	}*/

	return nil
}

func ChunkHash(x, z int32) uint64 {
	return uint64(uint32(z))<<32 | uint64(uint32(x))
}

func locationEntryToPos(index int) (x, z int32) {
	index = (index / 4) / 32

	return 0, 0
}
