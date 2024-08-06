package region

import (
	"bytes"
	//"compress/zlib"
	"encoding/binary"
	"os"

	"github.com/4kills/go-zlib"
	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/net/io/util"
)

// Encode writes the region file to w.
func (f *File) Encode(w *os.File) error {
	var locationTable [4096]byte
	var timestampTable [4096]byte
	var chunksOffset = len(locationTable) + len(timestampTable)

	var buf = new(bytes.Buffer)

	var chunkBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(chunkBuffer)

	for _, chunk := range f.chunks {
		locationIndex := ((uint32(chunk.X) % 32) + (uint32(chunk.Z)%32)*32) * 4

		binary.BigEndian.PutUint32(timestampTable[locationIndex:locationIndex+4], uint32(chunk.LastModified))

		offset := (buf.Len() + chunksOffset) / 4096

		locationTable[locationIndex+0],
			locationTable[locationIndex+1],
			locationTable[locationIndex+2],
			locationTable[locationIndex+3] =
			byte(offset>>16), byte(offset>>8), byte(offset), 1

		chunkBuffer.Reset()

		zlib := easyZlib{zlib.NewWriter(chunkBuffer)}

		f := util.NewFlusher(zlib)
		if err := nbt.NewEncoder(f).Encode("", chunkToAnvil(chunk)); err != nil {
			return err
		}
		if _, err := f.Flush(); err != nil {
			return err
		}

		if err := zlib.Close(); err != nil {
			return err
		}

		chunkLength := chunkBuffer.Len() + 1

		if _, err := buf.Write([]byte{
			byte(chunkLength >> 24),
			byte(chunkLength >> 16),
			byte(chunkLength >> 8),
			byte(chunkLength),
			2,
		}); err != nil {
			return err
		}

		if _, err := chunkBuffer.WriteTo(buf); err != nil {
			return err
		}

		if _, err := buf.Write(make([]byte, (4096-(buf.Len()%4096))%4096)); err != nil {
			return err
		}
	}

	if _, err := w.Write(locationTable[:]); err != nil {
		return err
	}
	if _, err := w.Write(timestampTable[:]); err != nil {
		return err
	}
	_, err := buf.WriteTo(w)

	return err
}

// easyzlib doesnt return an error if len(p) is 0
type easyZlib struct {
	*zlib.Writer
}

func (w easyZlib) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	return w.Writer.Write(p)
}
