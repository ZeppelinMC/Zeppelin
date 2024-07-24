package blocks

import (
	"bytes"
	_ "embed"

	"github.com/zeppelinmc/zeppelin/nbt"
)

type blockState struct {
	Id         int32             `json:"id"`
	Properties map[string]string `json:"properties"`
}

type Block struct {
	States []blockState `json:"states"`
}

func (b Block) FindState(properties map[string]string) (id int32, ok bool) {
	for _, state := range b.States {
		if isMapEqual(state.Properties, properties) {
			return state.Id, true
		}
	}
	return 0, false
}

func isMapEqual(m1, m2 map[string]string) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if len(m1) != len(m2) {
		return false
	}
	for key, value1 := range m1 {
		if m2[key] != value1 {
			return false
		}
	}

	return true
}

var Blocks map[string]Block

//go:embed blocks.nbt
var blockData []byte
var blockBuf = bytes.NewReader(blockData)

func LoadBlockCache() error {
	rd := nbt.NewStaticReader(blockBuf)
	_, _, _ = rd.ReadRoot(true)
	var compoundReader nbt.CompoundReader
	Blocks = make(map[string]Block)

	for {
		name, err, end := rd.Compound(&compoundReader)
		if end {
			break
		}
		if err != nil {
			return err
		}

		if err := compoundReader.ReadAll(func(len int32, rd nbt.ListReader) {
			states := make([]blockState, len)
			for i := int32(0); i < len; i++ {
				states[i].Properties = make(map[string]string)
				rd.Read([]any{&states[i].Id, states[i].Properties})

			}
			Blocks[name] = Block{
				States: states,
			}
		}); err != nil {
			return err
		}
	}

	return nil
}
