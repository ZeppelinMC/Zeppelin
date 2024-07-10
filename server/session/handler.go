package session

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
)

type handler func(*Session, packet.Packet)

// [state, packet id]
var Handlers = make(map[[2]int32]handler)

func (session *Session) handlePackets() {
	keepAlive := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-keepAlive.C:
			session.Conn.WritePacket(&play.ClientboundKeepAlive{KeepAliveID: time.Now().UnixMilli()})
		default:
			p, err := session.Conn.ReadPacket()
			if err != nil {
				log.Infof("[%s] Player %s disconnected\n", session.Conn.RemoteAddr(), session.Conn.Username())
				return
			}

			handler, ok := Handlers[[2]int32{session.Conn.State(), p.ID()}]
			if !ok {
				switch pk := p.(type) {
				case *configuration.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						session.clientName = string(pk.Data)
					}
				case *configuration.AcknowledgeFinishConfiguration:
					session.Conn.SetState(net.PlayState)

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
}
