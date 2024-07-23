package handler

import (
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatCommand, handleChatCommand)
	std.RegisterHandler(net.PlayState, play.PacketIdSignedChatCommand, handleChatCommand)
}

func handleChatCommand(session *std.StandardSession, p packet.Packet) {
	switch pk := p.(type) {
	case *play.ChatCommand:
		session.CommandManager().Call(pk.Command, session)
	}
}
