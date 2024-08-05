package region

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"os"

	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/net/buffers"
)

// Encode writes the region file to w. It uses a very high amount of memory so it should only be used when saving the world (stopping the server)
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

		zlib := zlib.NewWriter(chunkBuffer)

		if err := nbt.NewEncoder(zlib).Encode("", chunkToAnvil(chunk)); err != nil {
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
