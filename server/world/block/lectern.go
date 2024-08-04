package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Lectern struct {
	HasBook bool
	Powered bool
	Facing string
}

func (b Lectern) Encode() (string, BlockProperties) {
	return "minecraft:lectern", BlockProperties{
		"has_book": strconv.FormatBool(b.HasBook),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
	}
}

func (b Lectern) New(props BlockProperties) Block {
	return Lectern{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		HasBook: props["has_book"] != "false",
	}
}

func (b Lectern) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:lectern",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}