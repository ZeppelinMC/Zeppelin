package item

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

type Item struct {
	// The slot (as stored in the player data)
	Slot DataSlot `nbt:"Slot"`
	// the amount of items in the slot
	Count int32 `nbt:"count"`
	// The string id of this item
	Id string `nbt:"id"`
}
