package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type TrappedChest struct {
	Waterlogged bool
	Type string
	Facing string
}

func (b TrappedChest) Encode() (string, BlockProperties) {
	return "minecraft:trapped_chest", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
		"facing": b.Facing,
	}
}

func (b TrappedChest) New(props BlockProperties) Block {
	return TrappedChest{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
		Facing: props["facing"],
	}
}

func (b TrappedChest) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:trapped_chest",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}