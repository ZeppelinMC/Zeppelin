package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Beehive struct {
	Facing string
	HoneyLevel int
}

func (b Beehive) Encode() (string, BlockProperties) {
	return "minecraft:beehive", BlockProperties{
		"facing": b.Facing,
		"honey_level": strconv.Itoa(b.HoneyLevel),
	}
}

func (b Beehive) New(props BlockProperties) Block {
	return Beehive{
		Facing: props["facing"],
		HoneyLevel: atoi(props["honey_level"]),
	}
}

func (b Beehive) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:beehive",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}