package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Barrel struct {
	Facing string
	Open bool
}

func (b Barrel) Encode() (string, BlockProperties) {
	return "minecraft:barrel", BlockProperties{
		"facing": b.Facing,
		"open": strconv.FormatBool(b.Open),
	}
}

func (b Barrel) New(props BlockProperties) Block {
	return Barrel{
		Facing: props["facing"],
		Open: props["open"] != "false",
	}
}

func (b Barrel) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:barrel",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}