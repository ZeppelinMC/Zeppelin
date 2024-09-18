package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerPosition, handleMovement)
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerPositionAndRotation, handleMovement)
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerRotation, handleMovement)
	std.RegisterHandler(net.PlayState, play.PacketIdSetPlayerOnGround, handleMovement)
}

func handleMovement(s *std.StandardSession, p packet.Decodeable) {
	if s.AwaitingTeleportAcknowledgement.Load() {
		return
	}
	switch pk := p.(type) {
	case *play.SetPlayerPosition:
		s.Input.SetPosition(pk.X, pk.Y, pk.Z)
		s.Input.SetOnGround(pk.OnGround)
	case *play.SetPlayerPositionAndRotation:
		s.Input.SetPosition(pk.X, pk.Y, pk.Z)
		s.Input.SetRotation(pk.Yaw, pk.Pitch)
		s.Input.SetOnGround(pk.OnGround)
	case *play.SetPlayerRotation:
		s.Input.SetRotation(pk.Yaw, pk.Pitch)
		s.Input.SetOnGround(pk.OnGround)
	case *play.SetPlayerOnGround:
		s.Input.SetOnGround(pk.OnGround)
	}
}
