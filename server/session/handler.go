package session

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
)

type handler func(*Session, packet.Packet)

// id = LSB: state (1:configuration,0:play), packet id
var handlers = make(map[int32]handler)

func encodeHandlerID(conf bool, packetId int32) int32 {
	confi := int32(*(*byte)(unsafe.Pointer(&conf)))

	return packetId<<1 | confi
}

func RegisterHandler(state, packetId int32, handler handler) int {
	handlers[encodeHandlerID(state == net.ConfigurationState, packetId)] = handler

	return 0
}

func (session *Session) inConfiguration() bool {
	return session.Conn.State() == net.ConfigurationState
}

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

			handler, ok := handlers[encodeHandlerID(session.inConfiguration(), p.ID())]
			if !ok {
				switch pk := p.(type) {
				case *play.PlayerSession:
					session.HasSessionData.Set(true)
					session.SessionData.Set(*pk)
				case *configuration.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := io.ReadVarInt(pk.Data)
						session.clientName = string(data)
					}
				case *configuration.AcknowledgeFinishConfiguration:
					session.Conn.SetState(net.PlayState)

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
