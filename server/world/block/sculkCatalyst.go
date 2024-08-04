package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type SculkCatalyst struct {
	Bloom bool
}

func (b SculkCatalyst) Encode() (string, BlockProperties) {
	return "minecraft:sculk_catalyst", BlockProperties{
		"bloom": strconv.FormatBool(b.Bloom),
	}
}

func (b SculkCatalyst) New(props BlockProperties) Block {
	return SculkCatalyst{
		Bloom: props["bloom"] != "false",
	}
}

func (b SculkCatalyst) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:sculk_catalyst",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}