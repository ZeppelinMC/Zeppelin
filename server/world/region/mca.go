package region

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"sync"

	"github.com/zeppelinmc/zeppelin/nbt"
)

type RegionFile struct {
	reader io.ReaderAt

	locations []byte

	chunks map[int32]*Chunk
	chu_mu sync.Mutex
}

func chunkLocation(l int32) (offset, size int32) {
	offset = ((l >> 8) & 0xFFFFFF)
	size = l & 0xFF

	return offset * 4096, size * 4096
}

func (r *RegionFile) GetChunk(x, z int32) (*Chunk, error) {
	l := r.locations[((uint32(x)%32)+(uint32(z)%32)*32)*4:][:4]
	loc := int32(l[0])<<24 | int32(l[1])<<16 | int32(l[2])<<8 | int32(l[3])

	r.chu_mu.Lock()
	defer r.chu_mu.Unlock()
	if c, ok := r.chunks[loc]; ok {
		return c, nil
	}

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

	var data, _ = io.ReadAll(rd)

	r.chunks[loc] = &Chunk{}

	_, err = nbt.NewDecoder(bytes.NewReader(data)).Decode(r.chunks[loc])

	return r.chunks[loc], err
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
		chunks:    make(map[int32]*Chunk),
	}

	return nil
}
