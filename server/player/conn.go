package player

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/util/log"
	"time"
	"unsafe"
)

// registeredHandlers maps all handler functions by their state[0] and id[1]
var registeredHandlers = make(map[[2]int32]func(*Player, packet.Decodeable) error)

func RegisterHandler(state, id int32, handler func(*Player, packet.Decodeable) error) struct{} {
	registeredHandlers[[2]int32{state, id}] = handler

	return struct{}{}
}

// awaitPacket waits for the packet with the id to be received
func (p *Player) awaitPacket(id int32) (packet.Decodeable, error) {
	p.packetAwaited.Store(id)
	c := make(chan packet.Decodeable)
	p.packetAwaitedChan.Store(&c)

	pk := <-c
	if pk, ok := pk.(packet.Error); ok {
		return nil, pk.Error
	}

	return pk, nil
}

func interfaceAssert[T any](v any) T {
	return *(*T)(unsafe.Add(unsafe.Pointer(&v), unsafe.Sizeof(0)))
}

func (p *Player) keepAlive() error {
	l := time.Now().UnixMilli()
	p.cbLastKeepAlive.Store(l)
	return p.WritePacket(&play.ClientboundKeepAlive{KeepAliveID: l})
}

// isConnectionDead checks if more than 15 seconds have passed since the client sent a keep alive packet
func (p *Player) isConnectionDead() bool {
	lastKeepAliveByClient := p.sbLastKeepalive.Load()
	return lastKeepAliveByClient != 0 && time.Now().UnixMilli()-lastKeepAliveByClient > (15*1000)
}

func (p *Player) listenPackets() {
	keepAliveTicker := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-keepAliveTicker.C:
			if err := p.keepAlive(); err != nil {
				p.killConnection(false, "lost connection")
				return
			}
		default:
			if p.isConnectionDead() {
				//todo disconnect
				p.killConnection(false, "timed out")
				return
			}
			pk, s, err := p.ReadPacket()

			awaitId := p.packetAwaited.Load()

			if err != nil {
				p.killConnection(false, "lost connection")

				if awaitId != -1 {
					*p.packetAwaitedChan.Swap(nil) <- packet.Error{Error: err}
					p.packetAwaited.Store(-1)
				}
				return
			}

			if s {
				// The packet was stopped mid-interception, and shouldn't be handled
				continue
			}

			state, id := p.State(), pk.ID()

			// This packet is awaited by another goroutine
			if awaitId == id {
				*p.packetAwaitedChan.Swap(nil) <- pk
				p.packetAwaited.Store(-1)
			}

			handler, ok := registeredHandlers[[2]int32{state, id}]
			if !ok {
				switch pk := pk.(type) {
				case *configuration.ClientInformation, *play.ClientInformation:
					p.ClientInformation.Store(interfaceAssert[*configuration.ClientInformation](pk))
				}
				continue
			}

			if err := handler(p, pk); err != nil {
				p.killConnection(false, "packet handling error: "+err.Error())
				return
			}
		}
	}
}

// killConnection kills the player's connection
func (p *Player) killConnection(serverError bool, reason string) {
	var logFunction = log.Infolnf
	if serverError {
		// if an error happened on the server's side, the log message will be an error
		logFunction = log.Errorlnf
	}
	logFunction("%sPlayer %s disconnected: %s", log.FormatAddr(true /*TODO replace*/, p.RemoteAddr()), p.Username(), reason)
}
