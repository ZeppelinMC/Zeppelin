package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Conduit struct {
	Waterlogged bool
}

func (b Conduit) Encode() (string, BlockProperties) {
	return "minecraft:conduit", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Conduit) New(props BlockProperties) Block {
	return Conduit{
		Waterlogged: props["waterlogged"] != "false",
	}
}

func (b Conduit) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:conduit",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}