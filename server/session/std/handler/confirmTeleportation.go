package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdConfirmTeleportation, handleConfirmTeleportation)
}

func handleConfirmTeleportation(session *std.StandardSession, p packet.Decodeable) {
	session.AwaitingTeleportAcknowledgement.Store(false)
}
