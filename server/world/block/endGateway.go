package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type EndGateway struct {
}

func (b EndGateway) Encode() (string, BlockProperties) {
	return "minecraft:end_gateway", BlockProperties{}
}

func (b EndGateway) New(props BlockProperties) Block {
	return EndGateway{}
}

func (b EndGateway) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:end_gateway",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}