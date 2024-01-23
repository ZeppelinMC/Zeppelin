package anvil

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
	"strconv"
)

func WriteChunk(buf *bytes.Buffer, x, z int32, path string) error {
	chunkFile := "r." + strconv.FormatInt(int64(x>>5), 10) + "." + strconv.FormatInt(int64(z>>5), 10) + ".mca"

	f, err := os.OpenFile(path+chunkFile, os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	defer f.Close()

	offset, _, err := decodeChunkLocation(f, x, z)
	if err != nil {
		return err
	}

	_, compressionScheme, err := decodeChunkHeader(f, offset)
	if err != nil {
		return err
	}

	switch compressionScheme {
	//todo implement gzip decompression

	//zlib decompression
	case 2:
		//chunk header takes up 5 bytes
		_, _ = f.Seek(int64(offset+5), io.SeekStart)

		wr := zlib.NewWriter(f)

		_, _ = wr.Write(buf.Bytes())

		return wr.Close()
	}

	return nil
}
