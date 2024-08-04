package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type DecoratedPot struct {
	Cracked bool
	Facing string
	Waterlogged bool
}

func (b DecoratedPot) Encode() (string, BlockProperties) {
	return "minecraft:decorated_pot", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"cracked": strconv.FormatBool(b.Cracked),
		"facing": b.Facing,
	}
}

func (b DecoratedPot) New(props BlockProperties) Block {
	return DecoratedPot{
		Cracked: props["cracked"] != "false",
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}

func (b DecoratedPot) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:decorated_pot",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}