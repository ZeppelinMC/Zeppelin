package block

import (
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Jigsaw struct {
	Orientation string
}

func (b Jigsaw) Encode() (string, BlockProperties) {
	return "minecraft:jigsaw", BlockProperties{
		"orientation": b.Orientation,
	}
}

func (b Jigsaw) New(props BlockProperties) Block {
	return Jigsaw{
		Orientation: props["orientation"],
	}
}

func (b Jigsaw) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:jigsaw",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}