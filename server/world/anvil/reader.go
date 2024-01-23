package anvil

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"github.com/aimjel/minecraft/protocol"
	"github.com/aimjel/nitrate/server/world/chunk"
	"io"
	"os"
	"strconv"
)

type Reader struct {
	//path to the folder where the anvil files are stored
	path string
}

func NewReader(path string) *Reader {
	return &Reader{path: path}
}

func (r *Reader) ReadChunk(x, z int32) (*chunk.Chunk, error) {
	buf := protocol.GetBuffer(1024 * 30)
	defer protocol.PutBuffer(buf)

	if err := CopyChunk(buf, x, z, r.path); err != nil {
		return nil, fmt.Errorf("%v reading chunk %v %v", err, x, z)
	}

	return chunk.NewAnvilChunk(buf.Bytes())
}

func CopyChunk(buf *bytes.Buffer, x, z int32, path string) error {
	chunkFile := "r." + strconv.FormatInt(int64(x>>5), 10) + "." + strconv.FormatInt(int64(z>>5), 10) + ".mca"

	f, err := os.Open(path + chunkFile)
	if err != nil {
		return err
	}

	defer f.Close()

	offset, _, err := decodeChunkLocation(f, x, z)
	if err != nil {
		return err
	}

	chunkLength, compressionScheme, err := decodeChunkHeader(f, offset)
	if err != nil {
		return err
	}

	switch compressionScheme {
	//todo implement gzip decompression

	//zlib decompression
	case 2:
		//chunk header takes up 5 bytes
		_, _ = f.Seek(int64(offset+5), io.SeekStart)

		rd, err := zlib.NewReader(io.LimitReader(f, int64(chunkLength)))
		if err != nil {
			return err
		}

		if _, err = buf.ReadFrom(rd); err != nil {
			return err
		}
	}

	return nil
}

// decodeChunkLocation decodes the location entry for the x z Chunk coordinates passed.
// Returns the Chunk's offset in the file and sector or how much space it takes up.
func decodeChunkLocation(f *os.File, x, z int32) (uint32, uint32, error) {
	//allocate space for the location entry data
	loc := make([]byte, 4)

	//offset where the chunk location information starts
	offset := 4 * ((x & 31) + (z&31)*32)
	n, err := f.ReadAt(loc, int64(offset))
	if err != nil {
		return 0, 0, err
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

func decodeChunkHeader(f *os.File, offset uint32) (uint32, byte, error) {
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
