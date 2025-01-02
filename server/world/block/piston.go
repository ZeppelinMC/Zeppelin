package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Piston struct {
	Extended bool
	Facing string
}

func (b Piston) Encode() (string, BlockProperties) {
	return "minecraft:piston", BlockProperties{
		"extended": strconv.FormatBool(b.Extended),
		"facing": b.Facing,
	}
}

func (b Piston) New(props BlockProperties) Block {
	return Piston{
		Extended: props["extended"] != "false",
		Facing: props["facing"],
	}
}

func (b Piston) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:piston",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}