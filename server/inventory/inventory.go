package inventory

import (
	"sync"
	"unsafe"

	"slices"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"
)

type Inventory struct {
	sync.RWMutex
	inv []world.Slot
}

func From(inv []world.Slot) *Inventory {
	return &Inventory{inv: inv}
}

func (inv *Inventory) Data() []world.Slot {
	inv.RLock()
	defer inv.RUnlock()
	return inv.inv
}

func (inv *Inventory) Packet() (i []packet.Slot) {
	inv.Lock()
	defer inv.Unlock()
	i = make([]packet.Slot, 46)
	for _, slot := range inv.inv {
		item, ok := registry.GetItem(slot.Id)
		if !ok {
			continue
		}
		tag := *(*packet.SlotTag)(unsafe.Pointer(&slot.Tag))
		i[DataToNetwork(int(slot.Slot))] = packet.Slot{
			Present: true,
			Count:   slot.Count,
			Id:      item.ProtocolID,
			Tag:     tag,
		}
	}
	return
}

func (inv *Inventory) Slot(i int8) (world.Slot, bool) {
	inv.Lock()
	defer inv.Unlock()
	for _, s := range inv.inv {
		if s.Slot == i {
			return s, true
		}
	}
	return world.Slot{}, false
}

func (inv *Inventory) SetSlot(i int8, s world.Slot) {
	inv.Lock()
	defer inv.Unlock()
	for in, s := range inv.inv {
		if s.Slot == i {
			inv.inv[in] = s
			return
		}
	}
	inv.inv = append(inv.inv, s)
}

func (inv *Inventory) DeleteSlot(i int8) {
	inv.Lock()
	defer inv.Unlock()
	for in, s := range inv.inv {
		if s.Slot == i {
			inv.inv = slices.Delete(inv.inv, in, in+1)
			return
		}
	}
}

func DataToNetwork(index int) int {
	switch {
	case index == 100:
		index = 8
	case index == 101:
		index = 7
	case index == 102:
		index = 6
	case index == 103:
		index = 5
	case index == -106:
		index = 45
	case index <= 8:
		index += 36
	case index >= 80 && index <= 83:
		index -= 79
	}
	return index
}

func NetworkToData(index int) int {
	switch {
	case index == 8:
		index = 100
	case index == 7:
		index = 101
	case index == 6:
		index = 102
	case index == 5:
		index = 103
	case index == 45:
		index = -106
	case index >= 36 && index <= 43:
		index -= 36
	case index >= 1 && index <= 4:
		index += 79
	}
	return index
}
