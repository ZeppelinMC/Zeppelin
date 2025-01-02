package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type BlastFurnace struct {
	Facing string
	Lit bool
}

func (b BlastFurnace) Encode() (string, BlockProperties) {
	return "minecraft:blast_furnace", BlockProperties{
		"facing": b.Facing,
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b BlastFurnace) New(props BlockProperties) Block {
	return BlastFurnace{
		Lit: props["lit"] != "false",
		Facing: props["facing"],
	}
}

func (b BlastFurnace) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:blast_furnace",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}