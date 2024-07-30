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
	"github.com/zeppelinmc/zeppelin/server/world/region/section"
)

type RegionFile struct {
	reader io.ReaderAt

	locations []byte

	chunks map[uint64]*Chunk
	chu_mu sync.Mutex
}

func chunkLocation(l int32) (offset, size int32) {
	offset = ((l >> 8) & 0xFFFFFF)
	size = l & 0xFF

	return offset * 4096, size * 4096
}

func (r *RegionFile) GetChunk(x, z int32, generator Generator) (*Chunk, error) {
	hash := chunkHash(x, z)

	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	if c, ok := r.chunks[hash]; ok {
		return c, nil
	}

	c := generator.NewChunk(x, z)

	r.chunks[hash] = &c

	return &c, nil
	l := r.locations[((uint32(x)%32)+(uint32(z)%32)*32)*4:][:4]
	loc := int32(l[0])<<24 | int32(l[1])<<16 | int32(l[2])<<8 | int32(l[3])

	offset, size := chunkLocation(loc)
	if offset|size == 0 {
		return nil, fmt.Errorf("chunk %d %d not found", x, z)
	}

	var chunkHeader = make([]byte, 5)

	_, err := r.reader.ReadAt(chunkHeader, int64(offset))
	if err != nil {
		return nil, err
	}

	length := int32(chunkHeader[0])<<24 | int32(chunkHeader[1])<<16 | int32(chunkHeader[2])<<8 | int32(chunkHeader[3])
	compression := chunkHeader[4]

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

	var chunk anvilChunk

	_, err = nbt.NewDecoder(buf).Decode(&chunk)

	r.chunks[hash] = &Chunk{
		X:          chunk.XPos,
		Y:          chunk.YPos,
		Z:          chunk.ZPos,
		Heightmaps: chunk.Heightmaps,
	}

	r.chunks[hash].Sections = make([]*section.Section, len(chunk.Sections))
	for i, sec := range chunk.Sections {
		r.chunks[hash].Sections[i] = section.New(sec.Y, sec.BlockStates.Palette, sec.BlockStates.Data, sec.Biomes.Palette, sec.Biomes.Data, sec.SkyLight, sec.BlockLight)
	}

	return r.chunks[hash], err

	/*chunk, ok := r.chunks[loc]
	if !ok {
		return chunk, fmt.Errorf("not found chunk")
	}
	return chunk, nil*/
}

func DecodeRegion(r io.ReaderAt, f *RegionFile) error {
	var locationTable = make([]byte, 4096)

	_, err := r.ReadAt(locationTable, 0)
	if err != nil {
		return err
	}

	*f = RegionFile{
		reader: r,

		locations: locationTable,
		chunks:    make(map[uint64]*Chunk),
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

		f.chunks[loc] = &Chunk{}

		_, err = nbt.NewDecoder(chunkBuffer).Decode(f.chunks[loc])
		if err != nil {
			return err
		}
	}*/

	return nil
}

func chunkHash(x, z int32) uint64 {
	return uint64(x)<<32 | uint64(z)
}
