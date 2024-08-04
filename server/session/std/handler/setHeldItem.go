package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/text"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSetHeldItemServerbound, handleSetHeldItem)
}

func handleSetHeldItem(session *std.StandardSession, p packet.Packet) {
	switch pk := p.(type) {
	case *play.SetHeldItemServerbound:
		if pk.Slot < 0 || pk.Slot > 8 {
			session.Disconnect(text.TextComponent{Text: "Invalid slot"})
		}
		session.Player().SetSelectedItemSlot(int32(pk.Slot))
	}
}
