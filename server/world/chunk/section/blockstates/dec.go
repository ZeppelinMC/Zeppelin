// package blockstates provides parsing of the custom format for storing block states
package blockstates

import (
	"encoding/binary"
	"io"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/net/io/util"
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

		for i := 0; i < int(propertyCount); i++ {
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

// ReadHeader reads the header of the block states format. The header contains locations for each block
func ReadHeader(f io.Reader) (map[string]BlockLocation, error) {
	var blockCount int32
	if err := binary.Read(f, binary.BigEndian, &blockCount); err != nil {
		return nil, err
	}

	var locations = make(map[string]BlockLocation, blockCount)

	var strlen = make([]byte, 1)
	for i := 0; i < int(blockCount); i++ {
		name, err := readString(f, strlen)
		if err != nil {
			return nil, err
		}

		var location BlockLocation
		if err := binary.Read(f, binary.BigEndian, location); err != nil {
			return nil, err
		}

		locations[name] = location
	}

	return locations, nil
}

// reads a string prepended by a byte length, l should be a byte slice with a length of one, so make([]byte, 1) should work
func readString(f io.Reader, l []byte) (string, error) {
	if _, err := f.Read(l); err != nil {
		return "", err
	}
	var strdata = make([]byte, l[0])
	if _, err := f.Read(strdata); err != nil {
		return "", err
	}
	return unsafe.String(unsafe.SliceData(strdata), len(strdata)), nil
}
