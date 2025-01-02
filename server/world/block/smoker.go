package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Smoker struct {
	Facing string
	Lit bool
}

func (b Smoker) Encode() (string, BlockProperties) {
	return "minecraft:smoker", BlockProperties{
		"facing": b.Facing,
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b Smoker) New(props BlockProperties) Block {
	return Smoker{
		Facing: props["facing"],
		Lit: props["lit"] != "false",
	}
}

func (b Smoker) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:smoker",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}