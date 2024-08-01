package handler

import (
	"math"

	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/server/world/level"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdInteract, handleInteract)
}

func handleInteract(s *std.StandardSession, pk packet.Packet) {
	interact, ok := pk.(*play.Interact)
	if !ok {
		return
	}

	//TODO entity attack
	attacked, ok := s.Broadcast().SessionByEntityId(interact.EntityId)
	if !ok {
		return
	}
	if attacked.Player().Dimension() != s.Player().Dimension() {
		return
	}

	x1, y1, z1 := s.Player().Position()
	x2, y2, z2 := attacked.Player().Position()

	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))
	if distance > 4 {
		return
	}

	switch interact.Type {
	case play.InteractTypeAttack:
		if gm := attacked.Player().GameMode(); gm == level.GameModeCreative || gm == level.GameModeSpectator {
			return
		}
		s.Broadcast().DamageEvent(s, attacked, s.Player().Dimension(), "minecraft:player_attack")
	}
}
