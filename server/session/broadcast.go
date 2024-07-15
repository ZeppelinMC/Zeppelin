package session

import (
	"fmt"
	"math"
	"sync"

	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net/packet"
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

// Disconnects all the players on the broadcast
func (b *Broadcast) DisconnectAll(reason chat.TextComponent) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	for _, ses := range b.sessions {
		ses.Disconnect(reason)
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

	sesData, ok := session.SessionData()

	for _, ses := range b.sessions {
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

	id := session.Player().EntityId()

	for _, ses := range b.sessions {
		ses.PlayerInfoRemove(session.UUID())
		ses.DespawnEntities(id)
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

	session.PlayerInfoUpdate(toPlayerPk)
	b.sessions[session.UUID()] = session
}

func DegreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func (b *Broadcast) SpawnPlayer(session Session) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	x, y, z := session.Player().Position()
	yawdeg, pitchdeg := session.Player().Rotation()
	yaw, pitch := DegreesToAngle(yawdeg), DegreesToAngle(pitchdeg)

	spawnPacketSentToOthers := &play.SpawnEntity{
		EntityId:   session.Player().EntityId(),
		EntityUUID: session.UUID(),
		Type:       128, // player
		X:          x, Y: y, Z: z,
		Pitch: pitch,
		Yaw:   yaw,
	}
	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}

		x, y, z := ses.Player().Position()
		yawdeg, pitchdeg := ses.Player().Rotation()
		yaw, pitch := DegreesToAngle(yawdeg), DegreesToAngle(pitchdeg)

		ses.SpawnEntity(spawnPacketSentToOthers)
		session.SpawnEntity(&play.SpawnEntity{
			EntityId:   ses.Player().EntityId(),
			EntityUUID: ses.UUID(),
			Type:       128,
			X:          x,
			Y:          y,
			Z:          z,
			Yaw:        yaw,
			Pitch:      pitch,
		})
	}
}

func diffToBig(old, new float64) bool {
	a := new - old
	return a < -8 || a > 7.999755859375
}

// broadcasts the position and rotation changes to the server. should be used before setting the properties on the player
func (b *Broadcast) BroadcastPlayerMovement(session Session, x, y, z float64, yaw, pitch float32) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	var (
		oldX, oldY, oldZ = session.Player().Position()
		oldYaw, oldPitch = session.Player().Rotation()
	)

	eid := session.Player().EntityId()

	var pk packet.Packet
	switch {
	// changes in both position and rotation
	case (x != oldX || y != oldY || z != oldZ) && (yaw != oldYaw || pitch != oldPitch):
		yaw, pitch := DegreesToAngle(yaw), DegreesToAngle(pitch)
		pk = &play.UpdateEntityPositionAndRotation{
			EntityId: eid,
			DeltaX:   int16(x*4096 - oldX*4096),
			DeltaY:   int16(y*4096 - oldY*4096),
			DeltaZ:   int16(z*4096 - oldZ*4096),
			Yaw:      yaw,
			Pitch:    pitch,
		}
	case x != oldX || y != oldY || z != oldZ:
		pk = &play.UpdateEntityPosition{
			EntityId: eid,
			DeltaX:   int16(x*4096 - oldX*4096),
			DeltaY:   int16(y*4096 - oldY*4096),
			DeltaZ:   int16(z*4096 - oldZ*4096),
		}
	case yaw != oldYaw || pitch != oldPitch:
		yaw, pitch := DegreesToAngle(yaw), DegreesToAngle(pitch)
		pk = &play.UpdateEntityRotation{
			EntityId: eid,
			Yaw:      yaw,
			Pitch:    pitch,
		}
	}

	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		switch p := pk.(type) {
		case *play.UpdateEntityPosition:
			ses.UpdateEntityPosition(p)
		case *play.UpdateEntityPositionAndRotation:
			ses.UpdateEntityPositionRotation(p)
		case *play.UpdateEntityRotation:
			ses.UpdateEntityRotation(p)
		}
	}
}

func (b *Broadcast) Animation(session Session, animation byte) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	id := session.Player().EntityId()
	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		ses.EntityAnimation(id, animation)
	}
}

func (b *Broadcast) EntityMetadata(session Session, md map[byte]any) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	id := session.Player().EntityId()
	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		fmt.Println(ses.EntityMetadata(id, md))
	}
}
