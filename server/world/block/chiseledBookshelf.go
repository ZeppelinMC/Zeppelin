package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type ChiseledBookshelf struct {
	Slot4Occupied bool
	Slot5Occupied bool
	Facing string
	Slot0Occupied bool
	Slot1Occupied bool
	Slot2Occupied bool
	Slot3Occupied bool
}

func (b ChiseledBookshelf) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_bookshelf", BlockProperties{
		"slot_1_occupied": strconv.FormatBool(b.Slot1Occupied),
		"slot_2_occupied": strconv.FormatBool(b.Slot2Occupied),
		"slot_3_occupied": strconv.FormatBool(b.Slot3Occupied),
		"slot_4_occupied": strconv.FormatBool(b.Slot4Occupied),
		"slot_5_occupied": strconv.FormatBool(b.Slot5Occupied),
		"facing": b.Facing,
		"slot_0_occupied": strconv.FormatBool(b.Slot0Occupied),
	}
}

func (b ChiseledBookshelf) New(props BlockProperties) Block {
	return ChiseledBookshelf{
		Slot0Occupied: props["slot_0_occupied"] != "false",
		Slot1Occupied: props["slot_1_occupied"] != "false",
		Slot2Occupied: props["slot_2_occupied"] != "false",
		Slot3Occupied: props["slot_3_occupied"] != "false",
		Slot4Occupied: props["slot_4_occupied"] != "false",
		Slot5Occupied: props["slot_5_occupied"] != "false",
		Facing: props["facing"],
	}
}

func (b ChiseledBookshelf) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:chiseled_bookshelf",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}