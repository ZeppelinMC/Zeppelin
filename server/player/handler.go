package player

import (
	"aether/log"
	"aether/net"
	"aether/net/packet/configuration"
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
			player.clientInfo.Set(pk)
		case *configuration.ServerboundPluginMessage:
			if pk.Channel == "minecraft:brand" {
				player.clientName = string(pk.Data)
			}
		case *configuration.AcknowledgeFinishConfiguration:
			player.conn.SetState(net.PlayState)

			/*for x, z := int32(-6), int32(-6); x < 6 && z < 6; x, z = x+1, z+1 {
				c, _ := player.world.GetChunk(x, z)

				player.conn.WritePacket(c.Encode())
			}*/

			c, _ := player.world.GetChunk(0, 0)

			player.conn.WritePacket(c.Encode())

		default:
			fmt.Printf("unknown packet 0x%02x\n", p.ID())
		}
	}
}
