package player

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/util/log"
)

var _ = RegisterHandler(net.PlayState, play.PacketIdChatCommand, func(p *Player, packet packet.Decodeable) error {
	pk := packet.(*play.ChatCommand)

	log.Infolnf("%sPlayer %s (%s) issued server command: /%s", log.FormatAddr(p.serverProperties.LogIPs, p.RemoteAddr()), p.Username(), p.UUID(), pk.Command)
	p.commandManager.Call(pk.Command, p)

	return nil
})
