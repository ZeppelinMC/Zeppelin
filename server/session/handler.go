package session

import (
	"aether/log"
	"aether/net"
	"aether/net/packet"
	"aether/net/packet/configuration"
	"fmt"
	"runtime"
)

type handler func(*Session, packet.Packet)

// [state, packet id]
var Handlers = make(map[[2]int32]handler)

func (session *Session) handlePackets() {
	for {
		p, err := session.conn.ReadPacket()
		if err != nil {
			log.Infof("[%s] Player %s disconnected\n", session.conn.RemoteAddr(), session.conn.Username())
			return
		}

		handler, ok := Handlers[[2]int32{session.conn.State(), p.ID()}]
		if !ok {
			switch pk := p.(type) {
			case *configuration.ServerboundPluginMessage:
				if pk.Channel == "minecraft:brand" {
					session.clientName = string(pk.Data)
				}
			case *configuration.AcknowledgeFinishConfiguration:
				session.conn.SetState(net.PlayState)

				session.sendSpawnChunks()

				var stats runtime.MemStats
				runtime.ReadMemStats(&stats)

				fmt.Printf("Alloc: %dMiB, Total Alloc: %dMiB\n", stats.Alloc/1024/1024, stats.TotalAlloc/1024/1024)
			default:
				fmt.Sprintf("unknown packet 0x%02x\n", p.ID())
			}
			continue
		}
		handler(session, p)
	}
}
