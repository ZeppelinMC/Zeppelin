package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type EndPortal struct {
}

func (b EndPortal) Encode() (string, BlockProperties) {
	return "minecraft:end_portal", BlockProperties{}
}

func (b EndPortal) New(props BlockProperties) Block {
	return EndPortal{}
}

func (b EndPortal) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:end_portal",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}