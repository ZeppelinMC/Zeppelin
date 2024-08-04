package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Jukebox struct {
	HasRecord bool
}

func (b Jukebox) Encode() (string, BlockProperties) {
	return "minecraft:jukebox", BlockProperties{
		"has_record": strconv.FormatBool(b.HasRecord),
	}
}

func (b Jukebox) New(props BlockProperties) Block {
	return Jukebox{
		HasRecord: props["has_record"] != "false",
	}
}

func (b Jukebox) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:jukebox",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}