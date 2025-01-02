package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Furnace struct {
	Facing string
	Lit bool
}

func (b Furnace) Encode() (string, BlockProperties) {
	return "minecraft:furnace", BlockProperties{
		"facing": b.Facing,
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b Furnace) New(props BlockProperties) Block {
	return Furnace{
		Facing: props["facing"],
		Lit: props["lit"] != "false",
	}
}

func (b Furnace) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:furnace",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}