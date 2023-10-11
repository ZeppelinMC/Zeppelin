package anvil

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/dynamitemc/dynamite/server/world/chunk"
)

var buffers = sync.Pool{
	New: func() any {
		return bytes.NewBuffer(make([]byte, 0, 1024*10))
	},
}

type Reader struct {
	//path to the folder where the anvil files are stored
	path string
}

func NewReader(path string) *Reader {
	return &Reader{path: path}
}

func (r *Reader) ReadChunk(x, z int32) (*chunk.Chunk, error) {
	chunkFile := "r." + strconv.FormatInt(int64(x>>5), 10) + "." + strconv.FormatInt(int64(z>>5), 10) + ".mca"

	f, err := os.Open(r.path + chunkFile)
	if err != nil {
		return nil, fmt.Errorf("%v reading chunk %v %v", err, x, z)
	}

	defer f.Close()

	offset, _, err := r.decodeChunkLocation(f, x, z)
	if err != nil {
		return nil, err
	}

	chunkLength, compressionScheme, err := r.decodeChunkHeader(f, offset)
	if err != nil {
		return nil, err
	}

	//will hold the uncompressed nbt data
	buf := buffers.Get().(*bytes.Buffer)
	buf.Reset()

	switch compressionScheme {
	//todo implement gzip decompression

	//zlib decompression
	case 2:
		//todo handle error
		//chunk header takes up 5 bytes
		f.Seek(int64(offset+5), io.SeekStart)

		rd, err := zlib.NewReader(io.LimitReader(f, int64(chunkLength)))
		if err != nil {
			return nil, err
		}

		if _, err = buf.ReadFrom(rd); err != nil {
			return nil, err
		}
	}

	ch, err := chunk.NewAnvilChunk(buf.Bytes())
	buffers.Put(buf)
	return ch, err
}

// decodeChunkLocation decodes the location entry for the x z Chunk coordinates passed.
// Returns the Chunk's offset in the file and sector or how much space it takes up.
func (r *Reader) decodeChunkLocation(f *os.File, x, z int32) (uint32, uint32, error) {
	//allocate space for the location entry data
	loc := make([]byte, 4)

	//offset where the chunk location information starts
	offset := 4 * ((x & 31) + (z&31)*32)
	n, err := f.ReadAt(loc, int64(offset))
	if err != nil {
		return 0, 0, fmt.Errorf("%w reading chunk(x=%v z=%v) location", err, x, z)
	}

	if n != 4 {
		return 0, 0, fmt.Errorf("expected 4 bytes for location entry but got %v", n)
	}

	entry := binary.BigEndian.Uint32(loc)

	if entry == 0 {
		return 0, 0, chunk.ErrNotFound
	}

	return (entry >> 8) * 4096, (entry & 0xff) * 4096, nil
}

func (r *Reader) decodeChunkHeader(f *os.File, offset uint32) (uint32, byte, error) {
	//allocate space to store the header
	header := make([]byte, 5)

	n, err := f.ReadAt(header, int64(offset))
	if err != nil {
		return 0, 0, err
	}

	if n != 5 {
		return 0, 0, fmt.Errorf("expected 5 bytes for chunk header but got %v", n)
	}

	chunkLength := binary.BigEndian.Uint32(header[:4])
	compressionScheme := header[4]

	return chunkLength, compressionScheme, nil
}
