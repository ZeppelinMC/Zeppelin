package player

import (
	"aether/log"
	"aether/net"
	"aether/net/packet/configuration"
)

func (player *Player) handlePackets() {
	for {
		p, err := player.conn.ReadPacket()
		if err != nil {
			log.Infof("[%s] Player %s disconnected\n", player.conn.RemoteAddr(), player.conn.Username())
			return
		}

		switch pk := p.(type) {
		case *configuration.ClientInformation:
			player.clientInfo.Set(pk)
		case *configuration.ServerboundPluginMessage:
			if pk.Channel == "minecraft:brand" {
				player.clientName = string(pk.Data)
			}
		case *configuration.AcknowledgeFinishConfiguration:
			player.conn.SetState(net.PlayState)
		}
	}
}
