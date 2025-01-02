package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Dispenser struct {
	Facing string
	Triggered bool
}

func (b Dispenser) Encode() (string, BlockProperties) {
	return "minecraft:dispenser", BlockProperties{
		"triggered": strconv.FormatBool(b.Triggered),
		"facing": b.Facing,
	}
}

func (b Dispenser) New(props BlockProperties) Block {
	return Dispenser{
		Facing: props["facing"],
		Triggered: props["triggered"] != "false",
	}
}

func (b Dispenser) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:dispenser",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}