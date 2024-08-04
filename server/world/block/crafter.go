package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Crafter struct {
	Crafting bool
	Orientation string
	Triggered bool
}

func (b Crafter) Encode() (string, BlockProperties) {
	return "minecraft:crafter", BlockProperties{
		"crafting": strconv.FormatBool(b.Crafting),
		"orientation": b.Orientation,
		"triggered": strconv.FormatBool(b.Triggered),
	}
}

func (b Crafter) New(props BlockProperties) Block {
	return Crafter{
		Crafting: props["crafting"] != "false",
		Orientation: props["orientation"],
		Triggered: props["triggered"] != "false",
	}
}

func (b Crafter) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:crafter",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}