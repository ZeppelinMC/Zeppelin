package session

import (
	"sync"

	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/google/uuid"
)

type Broadcast struct {
	sessions    map[uuid.UUID]*Session
	sessions_mu sync.Mutex
}

func NewBroadcast() *Broadcast {
	return &Broadcast{
		sessions: make(map[uuid.UUID]*Session),
	}
}

func (b *Broadcast) addPlayer(session *Session) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	var toPlayerPk = &play.PlayerInfoUpdate{
		Actions: play.ActionAddPlayer | play.ActionUpdateListed | play.ActionInitializeChat,
		Players: map[uuid.UUID]play.PlayerAction{
			session.Conn.UUID(): {
				Name:       session.Conn.Username(),
				Properties: session.Conn.Properties(),

				Listed: true,
			},
		},
	}

	for _, ses := range b.sessions {
		ses.Conn.WritePacket(&play.PlayerInfoUpdate{
			Actions: play.ActionAddPlayer | play.ActionUpdateListed | play.ActionInitializeChat,
			Players: map[uuid.UUID]play.PlayerAction{
				session.Conn.UUID(): {
					Name:       session.Conn.Username(),
					Properties: session.Conn.Properties(),
					Listed:     true,
				},
			},
		})
		toPlayerPk.Players[ses.Conn.UUID()] = play.PlayerAction{
			Name:             ses.Conn.Username(),
			Properties:       ses.Conn.Properties(),
			Listed:           true,
			HasSignatureData: ses.HasSessionData.Get(),
			Session:          ses.SessionData.Get(),
		}
	}

	session.Conn.WritePacket(toPlayerPk)
	b.sessions[session.Conn.UUID()] = session
}
