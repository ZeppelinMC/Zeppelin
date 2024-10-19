package handler

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/server/world/level/item"

	slot0 "github.com/zeppelinmc/zeppelin/protocol/net/slot"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdClickContainer, handleClickContainer)
}

func handleClickContainer(s *std.StandardSession, pk packet.Decodeable) {
	ccpk, ok := pk.(*play.ClickContainer)
	if !ok {
		return
	}

	if s.WindowView.Load() != int32(ccpk.WindowId) {
		return // messing with a container other than the one open - ignore
	}

	if s.StateID.Load() != ccpk.State {
		switch ccpk.WindowId {
		case 0: // inventory
			s.SendInventory()
		default:
			_, w, ok := s.Dimension().WindowManager.Get(int32(ccpk.WindowId))
			if !ok {
				return // didnt find the window ////////////
			}
			s.SetContainerContent(*w)
		}
		return // synchronise container with the client
	}

	var w *container.Container

	switch ccpk.WindowId {
	case 0:
		w = s.Player().Inventory()
	default:
		_, c, ok := s.Dimension().WindowManager.Get(int32(ccpk.WindowId))
		if !ok {
			return
		}
		w = &c.Items
	}

	switch ccpk.Mode {
	case 0:
		switch ccpk.Button {
		case 0: // left click
			if ccpk.Slot == -999 {
				panic("unimplemented left-click-outside-inv")
			}

			slot := item.DataSlotFrom(int32(ccpk.Slot))

			fmt.Println(slot, ccpk.ChangedSlots, ccpk.CarriedItem)

			if !ccpk.CarriedItem.Is(slot0.Air) {
				s.CarriedItem, _ = w.Slot(slot)

				w.SetAt(item.Air, slot)
			} else {
				w.Merge(slot, &s.CarriedItem)
			}
		}
	}
}
