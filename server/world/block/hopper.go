package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Hopper struct {
	Enabled bool
	Facing string
}

func (b Hopper) Encode() (string, BlockProperties) {
	return "minecraft:hopper", BlockProperties{
		"enabled": strconv.FormatBool(b.Enabled),
		"facing": b.Facing,
	}
}

func (b Hopper) New(props BlockProperties) Block {
	return Hopper{
		Enabled: props["enabled"] != "false",
		Facing: props["facing"],
	}
}

func (b Hopper) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:hopper",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}