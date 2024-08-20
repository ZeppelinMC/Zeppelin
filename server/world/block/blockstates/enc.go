package blockstates

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/log"
)

func Encode(f *os.File, blocks map[string]Block) error {
	if _, err := f.Write(magic[:]); err != nil {
		return err
	}

	var buf = new(bytes.Buffer)

	var locations = make(map[string]BlockLocation, len(blocks))

	for name, block := range blocks {
		var loc = BlockLocation{Offset: int32(buf.Len())}

		if err := binary.Write(buf, binary.BigEndian, uint16(len(block))); err != nil {
			return err
		}

		for _, state := range block {
			if err := binary.Write(buf, binary.BigEndian, state.Id); err != nil {
				return err
			}
			if err := buf.WriteByte(byte(len(state.Properties))); err != nil {
				return err
			}
			for name, value := range state.Properties {
				if err := writeString(buf, name); err != nil {
					return err
				}
				if err := writeString(buf, value); err != nil {
					return err
				}
			}
		}

		loc.Size = int32(buf.Len()) - loc.Offset

		if name == "minecraft:stone" {
			log.Println(loc.Offset, loc.Size, buf.Bytes()[loc.Offset:loc.Offset+loc.Size])
		}

		locations[name] = loc
	}

	i32 := calcheaderlen(locations)

	for i, loc := range locations {
		loc.Offset += i32
		locations[i] = loc
	}

	_, err := writeLocations(f, locations)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(f)

	return err
}

func calcheaderlen(loc map[string]BlockLocation) int32 {
	i := 8

	for name := range loc {
		i += 1 + len(name) + 8
	}

	return int32(i)
}

func writeLocations(f io.Writer, locs map[string]BlockLocation) (i int, err error) {
	if err := binary.Write(f, binary.BigEndian, int32(len(locs))); err != nil {
		return i, err
	}
	i += 4
	for name, loc := range locs {
		if err := writeString(f, name); err != nil {
			return i, err
		}
		i += 1 + len(name)
		if err := binary.Write(f, binary.BigEndian, loc); err != nil {
			return i, err
		}
		i += 8
	}

	return i, nil
}

func writeString(f io.Writer, str string) error {
	if _, err := f.Write([]byte{byte(len(str))}); err != nil {
		return err
	}
	_, err := f.Write(unsafe.Slice(unsafe.StringData(str), len(str)))

	return err
}
