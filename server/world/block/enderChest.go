package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type EnderChest struct {
	Facing string
	Waterlogged bool
}

func (b EnderChest) Encode() (string, BlockProperties) {
	return "minecraft:ender_chest", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b EnderChest) New(props BlockProperties) Block {
	return EnderChest{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}

func (b EnderChest) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:ender_chest",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}