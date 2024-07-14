package std

import (
	"fmt"
	"runtime"
	"time"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
)

type handler func(*StandardSession, packet.Packet)

var handlers = make(map[[2]int32]handler)

func RegisterHandler(state, id int32, handler handler) {
	handlers[[2]int32{state, id}] = handler
}

func (session *StandardSession) inConfiguration() bool {
	return session.conn.State() == net.ConfigurationState
}

func (session *StandardSession) handlePackets() {
	keepAlive := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-keepAlive.C:
			session.conn.WritePacket(&play.ClientboundKeepAlive{KeepAliveID: time.Now().UnixMilli()})
		default:
			p, err := session.conn.ReadPacket()
			if err != nil {
				log.Infof("[%s] Player %s disconnected\n", session.conn.RemoteAddr(), session.conn.Username())
				session.broadcast.RemovePlayer(session)
				return
			}

			handler, ok := handlers[[2]int32{session.conn.State(), p.ID()}]
			if !ok {
				switch pk := p.(type) {
				case *play.PlayerSession:
					session.hasSessionData.Set(true)
					session.sessionData.Set(*pk)

					session.broadcast.UpdateSession(session)
				case *configuration.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := io.ReadVarInt(pk.Data)
						session.clientName = string(data)
					}
				case *configuration.AcknowledgeFinishConfiguration:
					session.conn.SetState(net.PlayState)

					session.sendSpawnChunks()

					var stats runtime.MemStats
					runtime.ReadMemStats(&stats)

					fmt.Printf("Alloc: %dMiB, Total Alloc: %dMiB\n", stats.Alloc/1024/1024, stats.TotalAlloc/1024/1024)
				default:
					fmt.Printf("unknown packet 0x%02x\n", p.ID())
				}
				continue
			}
			handler(session, p)
		}
	}
}
