package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Bell struct {
	Facing string
	Powered bool
	Attachment string
}

func (b Bell) Encode() (string, BlockProperties) {
	return "minecraft:bell", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"attachment": b.Attachment,
		"facing": b.Facing,
	}
}

func (b Bell) New(props BlockProperties) Block {
	return Bell{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Attachment: props["attachment"],
	}
}

func (b Bell) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:bell",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}