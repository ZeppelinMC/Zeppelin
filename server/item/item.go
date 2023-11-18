package item

import (
	"errors"
	"unsafe"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/registry"
)

var ErrInvalidID = errors.New("invalid id")

func hasTag(tag Tag) bool {
	return tag.Damage != 0 && tag.RepairCost != 0 && len(tag.Enchantments) != 0
}

func Is(item Item, item1 Item) bool {
	return item.Id == item1.Id && !hasTag(item.Tag) && !hasTag(item1.Tag)
}

var Air = Item{}

type Item struct {
	Count int8   `nbt:"Count"`
	Slot  int8   `nbt:"Slot"`
	Id    string `nbt:"id"`
	Tag   Tag    `nbt:"tag"`
}

type Enchantment struct {
	Id    string `nbt:"id"`
	Level int16  `nbt:"lvl"`
}

type Tag struct {
	Damage       int32         `nbt:"Damage"`
	RepairCost   int32         `nbt:"RepairCost"`
	Enchantments []Enchantment `nbt:"Enchantments"`
}

func (item Item) ToPacketSlot() (packet.Slot, error) {
	return ItemToPacketSlot(item)
}

func ItemToPacketSlot(item Item) (packet.Slot, error) {
	it, ok := registry.Item.Get(item.Id)
	if !ok {
		return packet.Slot{}, ErrInvalidID
	}
	return packet.Slot{
		Present: true,
		Count:   item.Count,
		Id:      it,
		Tag:     *(*packet.SlotTag)(unsafe.Pointer(&item.Tag)),
	}, nil
}

func PacketSlotToItem(sl int8, slot packet.Slot) (Item, error) {
	id, ok := registry.Item.Find(slot.Id)
	if !ok {
		return Item{}, ErrInvalidID
	}
	return Item{
		Count: slot.Count,
		Slot:  sl,
		Id:    id,
		Tag:   *(*Tag)(unsafe.Pointer(&slot.Tag)),
	}, nil
}
