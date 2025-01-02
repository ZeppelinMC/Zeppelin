package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type CommandBlock struct {
	Conditional bool
	Facing string
}

func (b CommandBlock) Encode() (string, BlockProperties) {
	return "minecraft:command_block", BlockProperties{
		"conditional": strconv.FormatBool(b.Conditional),
		"facing": b.Facing,
	}
}

func (b CommandBlock) New(props BlockProperties) Block {
	return CommandBlock{
		Conditional: props["conditional"] != "false",
		Facing: props["facing"],
	}
}

func (b CommandBlock) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:command_block",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}