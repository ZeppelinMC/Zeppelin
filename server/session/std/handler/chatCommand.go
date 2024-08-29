package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatCommand, handleChatCommand)
	std.RegisterHandler(net.PlayState, play.PacketIdSignedChatCommand, handleChatCommand)
}

func handleChatCommand(session *std.StandardSession, p packet.Decodeable) {
	switch pk := p.(type) {
	case *play.ChatCommand:
		session.CommandManager().Call(pk.Command, session)
	}
}
