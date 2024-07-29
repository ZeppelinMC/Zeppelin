package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/item"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/text"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSetCreativeModeSlot, handleSetCreativeSlot)
}

func handleSetCreativeSlot(s *std.StandardSession, pk packet.Packet) {
	scs, ok := pk.(*play.SetCreativeModeSlot)
	if !ok {
		return
	}
	// out of bounds
	if scs.Slot < 0 || scs.Slot > 45 {
		s.Disconnect(text.Unmarshalf(s.Config().Chat.Formatter.Rune(), "Creative mode slot out of bounds (expected 0<slot<46, got %d)", scs.Slot))
		return
	}
	gameMode := s.Player().GameMode()
	if gameMode != world.GameTypeCreative {
		s.Disconnect(text.TextComponent{Text: "Use of creative mode slot on a game mode other than creative is not allowed"})
		return
	}
	inv := s.Player().Inventory()

	item, err := item.New(int32(scs.Slot), scs.ClickedItem)
	if err != nil {
		s.Disconnect(text.Unmarshalf(s.Config().Chat.Formatter.Rune(), "Invalid item id %d", scs.ClickedItem.ItemId))
		return
	}

	inv.SetSlot(item)
}
