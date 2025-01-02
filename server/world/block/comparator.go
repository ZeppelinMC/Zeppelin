package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Comparator struct {
	Facing string
	Mode string
	Powered bool
}

func (b Comparator) Encode() (string, BlockProperties) {
	return "minecraft:comparator", BlockProperties{
		"facing": b.Facing,
		"mode": b.Mode,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b Comparator) New(props BlockProperties) Block {
	return Comparator{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		Mode: props["mode"],
	}
}

func (b Comparator) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:comparator",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}