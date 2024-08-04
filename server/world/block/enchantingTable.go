package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type EnchantingTable struct {
}

func (b EnchantingTable) Encode() (string, BlockProperties) {
	return "minecraft:enchanting_table", BlockProperties{}
}

func (b EnchantingTable) New(props BlockProperties) Block {
	return EnchantingTable{}
}

func (b EnchantingTable) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:enchanting_table",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}