package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdConfirmTeleportation, handleConfirmTeleportation)
}

func handleConfirmTeleportation(session *std.StandardSession, p packet.Packet) {
	session.AwaitingTeleportAcknowledgement.Set(false)

}
