package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type SculkShrieker struct {
	CanSummon bool
	Shrieking bool
	Waterlogged bool
}

func (b SculkShrieker) Encode() (string, BlockProperties) {
	return "minecraft:sculk_shrieker", BlockProperties{
		"can_summon": strconv.FormatBool(b.CanSummon),
		"shrieking": strconv.FormatBool(b.Shrieking),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SculkShrieker) New(props BlockProperties) Block {
	return SculkShrieker{
		CanSummon: props["can_summon"] != "false",
		Shrieking: props["shrieking"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}

func (b SculkShrieker) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:sculk_shrieker",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}