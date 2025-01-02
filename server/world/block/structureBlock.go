package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type StructureBlock struct {
	Mode string
}

func (b StructureBlock) Encode() (string, BlockProperties) {
	return "minecraft:structure_block", BlockProperties{
		"mode": b.Mode,
	}
}

func (b StructureBlock) New(props BlockProperties) Block {
	return StructureBlock{
		Mode: props["mode"],
	}
}

func (b StructureBlock) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:structure_block",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}