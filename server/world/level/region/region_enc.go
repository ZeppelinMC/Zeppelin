package region

import (
	"bytes"
	"fmt"
	"time"

	"encoding/binary"
	"os"

	"github.com/zeppelinmc/zeppelin/protocol/nbt"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/compress"
)

type bufferCloser struct {
	*bytes.Buffer
}

func (bufferCloser) Close() error {
	return nil
}

const (
	CompressionGzip = iota + 1
	CompressionZlib
	CompressionNone
	CompressionLZ4
)

// 4096B
var MaxCompressedPacketSize = 4096

// Encode writes itself to w.
func (f *File) Encode(w *os.File, compressionScheme byte) error {
	var locationTable [4096]byte
	var timestampTable [4096]byte
	var chunksOffset = len(locationTable) + len(timestampTable)

	var buf = new(bytes.Buffer)

	var chunkBuffer = buffers.Buffers.Get().(*bytes.Buffer)
	defer buffers.Buffers.Put(chunkBuffer)

	for _, chunk := range f.chunks {
		locationIndex := ((uint32(chunk.X) % 32) + (uint32(chunk.Z)%32)*32) * 4

		lastModified := time.Now().UnixMilli() // this has no purpose
		binary.BigEndian.PutUint32(timestampTable[locationIndex:locationIndex+4], uint32(lastModified))

		offset := (buf.Len() + chunksOffset) / 4096

		locationTable[locationIndex+0],
			locationTable[locationIndex+1],
			locationTable[locationIndex+2],
			locationTable[locationIndex+3] =
			byte(offset>>16), byte(offset>>8), byte(offset), 1

		chunkBuffer.Reset()

		if err := nbt.NewEncoder(chunkBuffer).Encode("", chunkToAnvil(chunk)); err != nil {
			return err
		}

		var data []byte
		var err error

		switch compressionScheme {
		case CompressionGzip:
			data, err = compress.CompressGzip(chunkBuffer.Bytes(), nil)
			if err != nil {
				return err
			}
		case CompressionZlib:
			data, err = compress.CompressZlib(chunkBuffer.Bytes(), nil)
			if err != nil {
				return err
			}
		case CompressionNone:
			data = chunkBuffer.Bytes()
		default:
			return fmt.Errorf("unknown compression scheme %d", compressionScheme)
		}

		chunkLength := chunkBuffer.Len() + 1

		if _, err := buf.Write([]byte{
			byte(chunkLength >> 24),
			byte(chunkLength >> 16),
			byte(chunkLength >> 8),
			byte(chunkLength),
			compressionScheme,
		}); err != nil {
			return err
		}

		if _, err := buf.Write(data); err != nil {
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
