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

// kills the session's ticker, and removes it from the server broadcast
func (session *StandardSession) kill(err bool, reason string) {
	var logfn = log.Infolnf
	if err {
		logfn = log.Errorlnf
	}
	logfn("%sPlayer %s disconnected: %s", log.FormatAddr(session.config.LogIPs, session.Addr()), session.Username(), reason)
	session.stopTick <- struct{}{}
	session.broadcast.RemovePlayer(session)
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
				session.kill(false, "timed out")
				return
			}
			p, s, err := session.conn.ReadPacket()
			if err != nil {
				session.kill(false, "lost connection")
				return
			}

			if s || session.readIntercept(p) {
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
						if err := session.updateBrand(pk.Data); err != nil {
							session.kill(true, "malformed brand")
						}
					}
				case *play.ServerboundPluginMessage:
					if pk.Channel == "minecraft:brand" {
						if err := session.updateBrand(pk.Data); err != nil {
							session.kill(true, "malformed brand")
						}
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

func (session *StandardSession) updateBrand(data []byte) (err error) {
	session.clientName, err = encoding.String(data)
	return
}
