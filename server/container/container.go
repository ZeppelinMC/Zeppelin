package container

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/protocol/net/slot"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/level/item"
)

// A container that holds items
type Container []item.Item

// NetworkConvert encodes the container with the specified size and changes the slot from data slots to network slots. This should be used for inventories
func (c Container) EncodeResize(size int) []slot.Slot {
	s := make([]slot.Slot, size)
	for _, item := range c {
		id, ok := registry.Item.Lookup(item.Id)
		if !ok {
			continue
		}
		s[item.Slot.Network()] = slot.Slot{
			ItemCount: item.Count,
			ItemId:    id,
		}
	}

	return s
}

func (c Container) Encode() []slot.Slot {
	s := make([]slot.Slot, len(c))
	for _, item := range c {
		id, ok := registry.Item.Lookup(item.Id)
		if !ok {
			continue
		}
		s[item.Slot] = slot.Slot{
			ItemCount: item.Count,
			ItemId:    id,
		}
	}

	return s
}

// Set adds the item to the container and replaces the existing one if found, and returns if the operation was successful
func (c *Container) Set(item item.Item) {
	for x, i := range *c {
		if i.Slot == item.Slot {
			(*c)[x] = item
			return
		}
	}
	*c = append(*c, item)
}

// Add adds the item to the container and replaces the existing one if found, and returns if the operation was successful
func (c *Container) SetAt(item item.Item, slot item.DataSlot) {
	item.Slot = slot
	c.Set(item)
}

// finds the item at the specified data slot
func (c Container) Slot(slot item.DataSlot) (item.Item, bool) {
	for _, item := range c {
		if item.Slot == slot {
			return item, true
		}
	}
	return item.Air, false
}

func (c *Container) Merge(slot item.DataSlot, carriedItem *item.Item) {
	if carriedItem.Is(item.Air) {
		return
	}
	it, ok := c.Slot(slot)
	if !ok {
		fmt.Println("nothing occupies slot! setting to", slot, *carriedItem)
		c.SetAt(*carriedItem, slot)

		s, _ := c.Slot(slot)
		fmt.Println("new item:", slot, s)
		return
	}
	if !carriedItem.Is(it) || it.Count > 64 {
		fmt.Println("swap", it, "with", *carriedItem)
		c.SetAt(*carriedItem, slot)
		*carriedItem = it

		s, _ := c.Slot(slot)
		fmt.Println("new item:", slot, s)
		return
	}

	it.Count += carriedItem.Count
	carriedItem.Count = 0

	if it.Count > 64 {
		carriedItem.Count = it.Count - 64

		it.Count = 64
	}

	if carriedItem.Count == 0 {
		*carriedItem = item.Air
	}

	fmt.Println("new item:", it)

	c.SetAt(it, slot)
}
