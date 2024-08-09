package std

import (
	"time"

	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/text"
)

var packetInterceptor func(s *StandardSession, pk packet.Packet, stop *bool)

type handler func(*StandardSession, packet.Packet)

var handlers = make(map[[2]int32]handler)

func RegisterHandler(state, id int32, handler handler) {
	handlers[[2]int32{state, id}] = handler
}

func SetPacketInterceptor(i func(s *StandardSession, pk packet.Packet, stop *bool)) {
	packetInterceptor = i
}

func (session *StandardSession) inConfiguration() bool {
	return session.conn.State() == net.ConfigurationState
}

func (session *StandardSession) intercept(pk packet.Packet) (stop bool) {
	if packetInterceptor != nil {
		packetInterceptor(session, pk, &stop)
	}

	return
}

func (session *StandardSession) handlePackets() {
	keepAlive := time.NewTicker(time.Second * 20)
	ticker := session.tick.New()
	for range ticker.C {
		select {
		case <-keepAlive.C:
			l := time.Now().UnixMilli()
			session.cbLastKeepAlive.Set(l)
			session.conn.WritePacket(&play.ClientboundKeepAlive{KeepAliveID: l})
		default:
			if lastKeepAlive := session.sbLastKeepalive.Get(); lastKeepAlive != 0 && time.Now().UnixMilli()-lastKeepAlive > (21*1000) {
				session.Disconnect(text.TextComponent{Text: "Timed out"})
			}
			p, err := session.conn.ReadPacket()
			if err != nil {
				session.broadcast.RemovePlayer(session)
				return
			}

			if session.intercept(p) {
				continue
			}

			handler, ok := handlers[[2]int32{session.conn.State(), p.ID()}]
			if !ok {
				switch pk := p.(type) {
				case *play.ChunkBatchReceived:
					session.awaitingChunkBatchAcknowledgement.Set(false)
				case *play.ServerboundKeepAlive:
					session.sbLastKeepalive.Set(time.Now().UnixMilli())
					session.broadcast.PlayerInfoUpdateLatency(session)
				case *play.PlayerSession:
					session.hasSessionData.Set(true)
					session.sessionData.Set(*pk)

					session.broadcast.PlayerInfoUpdateSession(session)
				case *configuration.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := io.ReadVarInt(pk.Data)
						session.clientName = string(data)
					}
				case *play.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := io.ReadVarInt(pk.Data)
						session.clientName = string(data)
					}
				case *configuration.AcknowledgeFinishConfiguration:
					session.conn.SetState(net.PlayState)
					session.login()
				default:
					log.Printlnf("Unknown packet 0x%02x %T", p.ID(), p)
				}
				continue
			}
			handler(session, p)
		}
	}
}
