package container

import (
	"github.com/zeppelinmc/zeppelin/net/slot"
	"github.com/zeppelinmc/zeppelin/server/item"
	"github.com/zeppelinmc/zeppelin/server/registry"
)

// A container that holds items
type Container []item.Item

func (c Container) Network() []slot.Slot {
	s := make([]slot.Slot, len(c))
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

func (c Container) Grow(newsize int) Container {
	if newsize <= len(c) {
		return c
	}
	newc := make(Container, newsize)
	copy(newc[:len(c)], c)

	return newc
}
