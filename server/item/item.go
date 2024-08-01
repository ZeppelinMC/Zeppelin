package item

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/net/slot"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
)

type DataSlot int8

func (slot DataSlot) Network() int32 {
	switch {
	case slot == 100:
		return 8
	case slot == 101:
		return 7
	case slot == 102:
		return 6
	case slot == 103:
		return 5
	case slot == -106:
		return 45
	case slot <= 8:
		return int32(slot + 36)
	case slot >= 80 && slot <= 83:
		return int32(slot - 79)
	default:
		return int32(slot)
	}
}

func DataSlotFrom(network int32) DataSlot {
	switch {
	case network == 8:
		return 100
	case network == 7:
		return 101
	case network == 6:
		return 102
	case network == 5:
		return 103
	case network == 45:
		return -106
	case network >= 36 && network <= 44:
		return DataSlot(network - 36)
	case network >= 1 && network <= 4:
		return DataSlot(network + 79)
	default:
		return DataSlot(network)
	}
}

type Item struct {
	// The slot (as stored in the player data)
	Slot DataSlot `nbt:"Slot"`
	// the amount of items in the slot
	Count int32 `nbt:"count"`
	// The string id of this item
	Id string `nbt:"id"`
	// Components of this item (https://minecraft.wiki/w/Data_component_format#List_of_components)
	Components struct {
		AttributeModifiers AttributeModifiers `nbt:"minecraft:attribute_modifiers"`
	} `nbt:"components"`
}

// returns the block of the item, if found
func (i Item) Block() (block section.Block, ok bool) {
	b := section.GetBlock(i.Id)
	_, ok = registry.Block.Lookup(i.Id)

	return b, ok
}

// New creates an item from the slot provided
func New(slot int32, item slot.Slot) (Item, error) {
	i := Item{
		Slot:  DataSlotFrom(slot),
		Count: item.ItemCount,
	}
	id, ok := registry.Item.NameOf(item.ItemId)
	if !ok {
		return i, fmt.Errorf("invalid item id")
	}

	i.Id = id

	return i, nil
}
