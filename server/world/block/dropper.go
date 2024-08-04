package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Dropper struct {
	Facing string
	Triggered bool
}

func (b Dropper) Encode() (string, BlockProperties) {
	return "minecraft:dropper", BlockProperties{
		"facing": b.Facing,
		"triggered": strconv.FormatBool(b.Triggered),
	}
}

func (b Dropper) New(props BlockProperties) Block {
	return Dropper{
		Triggered: props["triggered"] != "false",
		Facing: props["facing"],
	}
}

func (b Dropper) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:dropper",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}