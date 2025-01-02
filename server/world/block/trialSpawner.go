package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type TrialSpawner struct {
	TrialSpawnerState string
	Ominous bool
}

func (b TrialSpawner) Encode() (string, BlockProperties) {
	return "minecraft:trial_spawner", BlockProperties{
		"ominous": strconv.FormatBool(b.Ominous),
		"trial_spawner_state": b.TrialSpawnerState,
	}
}

func (b TrialSpawner) New(props BlockProperties) Block {
	return TrialSpawner{
		Ominous: props["ominous"] != "false",
		TrialSpawnerState: props["trial_spawner_state"],
	}
}

func (b TrialSpawner) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:trial_spawner",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}