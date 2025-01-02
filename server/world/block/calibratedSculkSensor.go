package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type CalibratedSculkSensor struct {
	Facing string
	Power int
	SculkSensorPhase string
	Waterlogged bool
}

func (b CalibratedSculkSensor) Encode() (string, BlockProperties) {
	return "minecraft:calibrated_sculk_sensor", BlockProperties{
		"facing": b.Facing,
		"power": strconv.Itoa(b.Power),
		"sculk_sensor_phase": b.SculkSensorPhase,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CalibratedSculkSensor) New(props BlockProperties) Block {
	return CalibratedSculkSensor{
		Facing: props["facing"],
		Power: atoi(props["power"]),
		SculkSensorPhase: props["sculk_sensor_phase"],
		Waterlogged: props["waterlogged"] != "false",
	}
}

func (b CalibratedSculkSensor) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:calibrated_sculk_sensor",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}