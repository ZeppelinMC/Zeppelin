package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/text"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdPlayerAbilitiesServerbound, handlePlayerAbilities)
}

func handlePlayerAbilities(s *std.StandardSession, pk packet.Packet) {
	abs, ok := pk.(*play.PlayerAbilitiesServerbound)
	if !ok {
		return
	}
	flying := abs.Flags&play.PlayerAbsFlying != 0
	abilities := s.Player().Abilities()
	if flying && !abilities.MayFly {
		s.Disconnect(text.TextComponent{Text: "You are not allowed to fly"})
		return
	}
	abilities.Flying = flying
	s.Player().SetAbilities(abilities)
}
