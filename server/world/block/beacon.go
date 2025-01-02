package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Beacon struct {
}

func (b Beacon) Encode() (string, BlockProperties) {
	return "minecraft:beacon", BlockProperties{}
}

func (b Beacon) New(props BlockProperties) Block {
	return Beacon{}
}

func (b Beacon) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:beacon",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}