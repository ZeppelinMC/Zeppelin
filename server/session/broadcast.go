package session

import (
	"sync"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/google/uuid"
)

type Broadcast struct {
	sessions    map[uuid.UUID]Session
	sessions_mu sync.Mutex
}

func NewBroadcast() *Broadcast {
	return &Broadcast{
		sessions: make(map[uuid.UUID]Session),
	}
}

// Returns a session by uuid
func (b *Broadcast) Session(uuid uuid.UUID) (ses Session, ok bool) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	ses, ok = b.sessions[uuid]

	return
}

// Returns a session by username
func (b *Broadcast) SessionByUsername(username string) (ses Session, ok bool) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, session := range b.sessions {
		if session.Username() == username {
			return session, true
		}
	}

	return nil, false
}

// Returns a session by entity id
func (b *Broadcast) SessionByEntityId(entityId int32) (ses Session, ok bool) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, session := range b.sessions {
		if session.Player().EntityId() == entityId {
			return session, true
		}
	}

	return nil, false
}

// when a player's session data updates
func (b *Broadcast) UpdateSession(session Session) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, ses := range b.sessions {
		sesData, ok := session.SessionData()
		ses.PlayerInfoUpdate(&play.PlayerInfoUpdate{
			Actions: play.ActionInitializeChat,
			Players: map[uuid.UUID]play.PlayerAction{
				session.UUID(): {
					HasSignatureData: ok,
					Session:          sesData,
				},
			},
		})
	}
}

// when a player leaves the server
func (b *Broadcast) RemovePlayer(session Session) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	delete(b.sessions, session.UUID())

	for _, ses := range b.sessions {
		ses.PlayerInfoRemove(session.UUID())
	}

}

// when a new player joins the server
func (b *Broadcast) AddPlayer(session Session) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	var toPlayerPk = &play.PlayerInfoUpdate{
		Actions: play.ActionAddPlayer | play.ActionUpdateListed | play.ActionInitializeChat,
		Players: map[uuid.UUID]play.PlayerAction{
			session.UUID(): {
				Name:       session.Username(),
				Properties: session.Properties(),

				Listed: true,
			},
		},
	}

	for _, ses := range b.sessions {
		ses.PlayerInfoUpdate(&play.PlayerInfoUpdate{
			Actions: play.ActionAddPlayer | play.ActionUpdateListed | play.ActionInitializeChat,
			Players: map[uuid.UUID]play.PlayerAction{
				session.UUID(): {
					Name:       session.Username(),
					Properties: session.Properties(),
					Listed:     true,
				},
			},
		})
		sesData, ok := ses.SessionData()
		toPlayerPk.Players[ses.UUID()] = play.PlayerAction{
			Name:             ses.Username(),
			Properties:       ses.Properties(),
			Listed:           true,
			HasSignatureData: ok,
			Session:          sesData,
		}
	}

	posX, posY, posZ := session.Player().Position()
	log.Infof("[%s] Player %s (%s) joined with entity id %d (%f %f %f)\n", session.Addr(), session.Username(), session.UUID(), session.Player().EntityId(), posX, posY, posZ)

	session.PlayerInfoUpdate(toPlayerPk)
	b.sessions[session.UUID()] = session
}
