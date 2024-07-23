package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdSwingArm, handleSwingArm)
}

func handleSwingArm(session *std.StandardSession, pk packet.Packet) {
	swing, ok := pk.(*play.SwingArm)
	if !ok {
		return
	}

	var id byte
	if swing.Hand == 1 {
		id = 3
	}

	session.Broadcast().Animation(session, id)
}
