package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type BrewingStand struct {
	HasBottle0 bool
	HasBottle1 bool
	HasBottle2 bool
}

func (b BrewingStand) Encode() (string, BlockProperties) {
	return "minecraft:brewing_stand", BlockProperties{
		"has_bottle_0": strconv.FormatBool(b.HasBottle0),
		"has_bottle_1": strconv.FormatBool(b.HasBottle1),
		"has_bottle_2": strconv.FormatBool(b.HasBottle2),
	}
}

func (b BrewingStand) New(props BlockProperties) Block {
	return BrewingStand{
		HasBottle0: props["has_bottle_0"] != "false",
		HasBottle1: props["has_bottle_1"] != "false",
		HasBottle2: props["has_bottle_2"] != "false",
	}
}

func (b BrewingStand) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:brewing_stand",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}