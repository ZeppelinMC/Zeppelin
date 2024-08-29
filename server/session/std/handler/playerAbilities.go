package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdPlayerAbilitiesServerbound, handlePlayerAbilities)
}

func handlePlayerAbilities(s *std.StandardSession, pk packet.Decodeable) {
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
