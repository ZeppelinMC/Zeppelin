// package blockstates provides parsing of the custom format for storing block states
package blockstates

import (
	"encoding/binary"
	"fmt"
	"io"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/util"
)

type BlockLocation struct {
	Offset, Size int32
}

type BlockState struct {
	Id         int32
	Properties map[string]string
}

type Block []BlockState

// ReadBlock reads the block states at the offset and size
func ReadBlock(f io.ReaderAt, loc BlockLocation) (Block, error) {
	var blockStates Block

	maxxer := util.NewReaderAtMaxxer(f, int(loc.Size), int64(loc.Offset))

	var stateCount uint16
	if err := binary.Read(maxxer, binary.BigEndian, &stateCount); err != nil {
		return blockStates, err
	}

	blockStates = make(Block, stateCount)

	var stringlen = make([]byte, 1)

	for i := range blockStates {
		var (
			blockStateId  int32
			propertyCount uint8
		)
		if err := binary.Read(maxxer, binary.BigEndian, &blockStateId); err != nil {
			return blockStates, err
		}
		if err := binary.Read(maxxer, binary.BigEndian, &propertyCount); err != nil {
			return blockStates, err
		}

		blockStates[i].Id = blockStateId
		blockStates[i].Properties = make(map[string]string, propertyCount)

		for j := 0; j < int(propertyCount); j++ {
			propertyName, err := readString(maxxer, stringlen)
			if err != nil {
				return blockStates, nil
			}
			propertyValue, err := readString(maxxer, stringlen)
			if err != nil {
				return blockStates, nil
			}
			blockStates[i].Properties[propertyName] = propertyValue
		}
	}

	return blockStates, nil
}

var magic = [...]byte{0x0F, 0x06, 0x60, 0xF0}

// ReadHeader reads the header of the block states format. The header contains locations for each block
func ReadHeader(f io.Reader) (map[string]BlockLocation, error) {
	var rmagic [4]byte
	if _, err := f.Read(rmagic[:]); err != nil {
		return nil, err
	}
	if rmagic != magic {
		return nil, fmt.Errorf("invalid magic header")
	}

	var locations_tmp [8]byte

	if _, err := f.Read(locations_tmp[:4]); err != nil {
		return nil, err
	}

	blockCount := int32(locations_tmp[0])<<24 | int32(locations_tmp[1])<<16 | int32(locations_tmp[2])<<8 | int32(locations_tmp[3])

	var locations = make(map[string]BlockLocation, blockCount)

	var strlen = make([]byte, 1)
	for i := 0; i < int(blockCount); i++ {
		name, err := readString(f, strlen)
		if err != nil {
			return nil, err
		}

		var location BlockLocation

		if _, err := f.Read(locations_tmp[:]); err != nil {
			return nil, err
		}
		location.Offset = int32(locations_tmp[0])<<24 | int32(locations_tmp[1])<<16 | int32(locations_tmp[2])<<8 | int32(locations_tmp[3])
		location.Size = int32(locations_tmp[4])<<24 | int32(locations_tmp[5])<<16 | int32(locations_tmp[6])<<8 | int32(locations_tmp[7])

		locations[name] = location
	}

	return locations, nil
}

// reads a string prepended by a byte length, l should be a byte slice with a length of one, so make([]byte, 1) will work. nil will also work
func readString(f io.Reader, l []byte) (string, error) {
	if l == nil {
		l = make([]byte, 1)
	}
	if _, err := f.Read(l); err != nil {
		return "", err
	}
	var strdata = make([]byte, l[0])
	if _, err := f.Read(strdata); err != nil {
		return "", err
	}
	return unsafe.String(unsafe.SliceData(strdata), len(strdata)), nil
}
