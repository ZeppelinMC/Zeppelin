package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSetHeldItemServerbound, handleSetHeldItem)
}

func handleSetHeldItem(session *std.StandardSession, p packet.Decodeable) {
	switch pk := p.(type) {
	case *play.SetHeldItemServerbound:
		if pk.Slot < 0 || pk.Slot > 8 {
			session.Disconnect(text.TextComponent{Text: "Invalid slot"})
		}
		session.Player().SetSelectedItemSlot(int32(pk.Slot))
	}
}
