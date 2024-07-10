package player

import (
	"aether/log"
	"aether/net"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"fmt"
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
			player.clientInfo.Set(*pk)
		case *play.ClientInformation:
			player.clientInfo.Set(pk.ClientInformation)
		case *configuration.ServerboundPluginMessage:
			if pk.Channel == "minecraft:brand" {
				player.clientName = string(pk.Data)
			}
		case *configuration.AcknowledgeFinishConfiguration:
			player.conn.SetState(net.PlayState)

			player.sendSpawnChunks()
		default:
			fmt.Sprintf("unknown packet 0x%02x\n", p.ID())
		}
	}
}
