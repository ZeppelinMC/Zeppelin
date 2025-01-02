package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type DaylightDetector struct {
	Power int
	Inverted bool
}

func (b DaylightDetector) Encode() (string, BlockProperties) {
	return "minecraft:daylight_detector", BlockProperties{
		"inverted": strconv.FormatBool(b.Inverted),
		"power": strconv.Itoa(b.Power),
	}
}

func (b DaylightDetector) New(props BlockProperties) Block {
	return DaylightDetector{
		Inverted: props["inverted"] != "false",
		Power: atoi(props["power"]),
	}
}

func (b DaylightDetector) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:daylight_detector",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}