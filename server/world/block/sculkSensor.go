package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type SculkSensor struct {
	Power int
	SculkSensorPhase string
	Waterlogged bool
}

func (b SculkSensor) Encode() (string, BlockProperties) {
	return "minecraft:sculk_sensor", BlockProperties{
		"power": strconv.Itoa(b.Power),
		"sculk_sensor_phase": b.SculkSensorPhase,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SculkSensor) New(props BlockProperties) Block {
	return SculkSensor{
		Power: atoi(props["power"]),
		SculkSensorPhase: props["sculk_sensor_phase"],
		Waterlogged: props["waterlogged"] != "false",
	}
}

func (b SculkSensor) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:sculk_sensor",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}