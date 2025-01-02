package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type ShulkerBox struct {
	Facing string
}

func (b ShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b ShulkerBox) New(props BlockProperties) Block {
	return ShulkerBox{
		Facing: props["facing"],
	}
}

func (b ShulkerBox) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:shulker_box",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}