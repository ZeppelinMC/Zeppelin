package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Campfire struct {
	SignalFire bool
	Waterlogged bool
	Facing string
	Lit bool
}

func (b Campfire) Encode() (string, BlockProperties) {
	return "minecraft:campfire", BlockProperties{
		"signal_fire": strconv.FormatBool(b.SignalFire),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b Campfire) New(props BlockProperties) Block {
	return Campfire{
		SignalFire: props["signal_fire"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
		Lit: props["lit"] != "false",
	}
}

func (b Campfire) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:campfire",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}