package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/util/log"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatCommand, handleChatCommand)
	std.RegisterHandler(net.PlayState, play.PacketIdSignedChatCommand, handleChatCommand)
}

func handleChatCommand(session *std.StandardSession, p packet.Decodeable) {
	switch pk := p.(type) {
	case *play.ChatCommand:
		log.Infolnf("%sPlayer %s (%s) issued server command: /%s", log.FormatAddr(session.Config().LogIPs, session.Addr()), session.Username(), session.UUID(), pk.Command)
		session.CommandManager().Call(pk.Command, session)
	}
}
