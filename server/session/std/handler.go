package std

import (
	"time"

	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/util/log"
)

var PacketReadInterceptor func(s *StandardSession, pk packet.Decodeable) bool
var PacketWriteInterceptor func(s *StandardSession, pk packet.Encodeable) bool

type handler func(*StandardSession, packet.Decodeable)

var handlers = make(map[[2]int32]handler)

func RegisterHandler(state, id int32, handler handler) {
	handlers[[2]int32{state, id}] = handler
}
func (session *StandardSession) inConfiguration() bool {
	return session.conn.State() == net.ConfigurationState
}

func (session *StandardSession) readIntercept(pk packet.Decodeable) (stop bool) {
	if PacketReadInterceptor != nil {
		return PacketReadInterceptor(session, pk)
	}

	return
}

func (session *StandardSession) handlePackets() {
	keepAlive := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-keepAlive.C:
			l := time.Now().UnixMilli()
			session.cbLastKeepAlive.Store(l)
			session.conn.WritePacket(&play.ClientboundKeepAlive{KeepAliveID: l})
		default:
			if lastKeepAlive := session.sbLastKeepalive.Load(); lastKeepAlive != 0 && time.Now().UnixMilli()-lastKeepAlive > (21*1000) {
				session.Disconnect(text.TextComponent{Text: "Timed out"})
			}
			p, err := session.conn.ReadPacket()
			if err != nil {
				log.Infolnf("[%s] Player %s disconnected: lost connection", session.Addr(), session.Username())
				session.broadcast.RemovePlayer(session)
				return
			}

			if session.readIntercept(p) {
				continue
			}

			handler, ok := handlers[[2]int32{session.conn.State(), p.ID()}]
			if !ok {
				switch pk := p.(type) {
				case *play.ChunkBatchReceived:
					session.chunksPerTick.Store(int32(pk.ChunksPerTick))
					session.awaitingChunkBatchAcknowledgement.Store(false)
				case *play.ServerboundKeepAlive:
					session.sbLastKeepalive.Store(time.Now().UnixMilli())
					session.broadcast.PlayerInfoUpdateLatency(session)
				case *play.PlayerSession:
					session.hasSessionData.Store(true)
					session.sessionData.Set(*pk)

					session.broadcast.PlayerInfoUpdateSession(session)
				case *configuration.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := encoding.VarInt(pk.Data)
						session.clientName = string(data)
					}
				case *play.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						_, data, _ := encoding.VarInt(pk.Data)
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
