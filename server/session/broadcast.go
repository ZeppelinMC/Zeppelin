package session

import (
	"math"
	"sync"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/net/packet/status"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/text"
	"github.com/zeppelinmc/zeppelin/util"
)

type Broadcast struct {
	sessions    map[uuid.UUID]Session
	sessions_mu sync.RWMutex

	console Session
}

func NewBroadcast(console Session) *Broadcast {
	return &Broadcast{
		sessions: make(map[uuid.UUID]Session),
		console:  console,
	}
}

// Disconnects all the players on the broadcast
func (b *Broadcast) DisconnectAll(reason text.TextComponent) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	for _, ses := range b.sessions {
		ses.Disconnect(reason)
	}
}

// Returns a session by uuid
func (b *Broadcast) Session(uuid uuid.UUID) (ses Session, ok bool) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	ses, ok = b.sessions[uuid]

	return
}

// Returns a session by uuid without locking the mutex
func (b *Broadcast) AsyncSession(uuid uuid.UUID) (ses Session, ok bool) {
	ses, ok = b.sessions[uuid]

	return
}

// Returns a session by username
func (b *Broadcast) SessionByUsername(username string) (ses Session, ok bool) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, session := range b.sessions {
		if session.Username() == username {
			return session, true
		}
	}

	return nil, false
}

// Returns a session by entity id
func (b *Broadcast) SessionByEntityId(entityId int32) (ses Session, ok bool) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, session := range b.sessions {
		if session.Player().EntityId() == entityId {
			return session, true
		}
	}

	return nil, false
}

// when a player's session data updates
func (b *Broadcast) UpdateSession(session Session) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

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

func (b *Broadcast) RemoveUUIDs(disconnectionReason text.TextComponent, uuids ...uuid.UUID) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()
	for _, uuid := range uuids {
		session, ok := b.sessions[uuid]
		if !ok {
			continue
		}
		session.Disconnect(disconnectionReason)
		delete(b.sessions, uuid)
	}
}

// when a player leaves the server
func (b *Broadcast) RemovePlayer(session Session) {
	log.Infolnf("[%s] Player %s disconnected", session.Addr(), session.Username())
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

func (b *Broadcast) SpawnPlayer(session Session) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	x, y, z := session.Player().Position()

	log.Infolnf("[%s] Player %s (%s) joined with entity id %d (%f %f %f)", session.Addr(), session.Username(), session.UUID(), session.Player().EntityId(), x, y, z)

	dim := session.Player().Dimension()

	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		if ses.Player().Dimension() != dim {
			continue
		}

		ses.SpawnPlayer(session)
		session.SpawnPlayer(ses)
	}
}

// broadcasts the position and rotation changes to the server. should be used before setting the properties on the player
func (b *Broadcast) BroadcastPlayerMovement(session Session, x, y, z float64, yaw, pitch float32) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	var (
		oldX, oldY, oldZ = session.Player().Position()
		oldYaw, oldPitch = session.Player().Rotation()
	)

	player := session.Player()

	dim := session.Player().Dimension()

	//fmt.Printf("%s new: %f %f %f, old: %f %f %f\n", session.Username(), x, y, z, oldX, oldY, oldZ)

	eid := session.Player().EntityId()

	var pk packet.Packet
	switch {
	// changes in both position and rotation
	case (x != oldX || y != oldY || z != oldZ) && (yaw != oldYaw || pitch != oldPitch):
		yaw, pitch := util.DegreesToAngle(yaw), util.DegreesToAngle(pitch)
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
		yaw, pitch := util.DegreesToAngle(yaw), util.DegreesToAngle(pitch)
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
		if ses.Player().Dimension() != dim {
			continue
		}
		switch p := pk.(type) {
		case *play.UpdateEntityPosition:
			ses.UpdateEntityPosition(player, p)
		case *play.UpdateEntityPositionAndRotation:
			ses.UpdateEntityPositionRotation(player, p)
		case *play.UpdateEntityRotation:
			ses.UpdateEntityRotation(player, p)
		}
	}
}

func (b *Broadcast) Animation(session Session, animation byte) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	id := session.Player().EntityId()
	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		ses.EntityAnimation(id, animation)
	}
}

func (b *Broadcast) EntityMetadata(session Session, md metadata.Metadata) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	id := session.Player().EntityId()
	for _, ses := range b.sessions {
		if ses.UUID() == session.UUID() {
			continue
		}
		ses.EntityMetadata(id, md)
	}
}

func (b *Broadcast) UpdateTimeForAll(worldAge, dayTime int64) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, ses := range b.sessions {
		ses.UpdateTime(worldAge, dayTime)
	}
}

func (b *Broadcast) BlockAction(x, y, z int32, dimension string, actionId, actionParameter byte) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	pk := &play.BlockAction{
		X: x, Y: y, Z: z,
		ActionId:        actionId,
		ActionParameter: actionParameter,
	}

	for _, ses := range b.sessions {
		if ses.Player().Dimension() != dimension {
			continue
		}
		ses.BlockAction(pk)
	}
}

func (b *Broadcast) PlaySound(pk *play.SoundEffect, dimension string) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, ses := range b.sessions {
		if ses.Player().Dimension() != dimension {
			continue
		}
		ses.PlaySound(pk)
	}
}

// this updates the block for everyone. It doesn't set the block in the world
func (b *Broadcast) UpdateBlock(x, y, z int32, block section.Block, dimension string) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, ses := range b.sessions {
		if ses.Player().Dimension() != dimension {
			continue
		}
		ses.UpdateBlock(x, y, z, block)
	}
}

// this updates the block entity for everyone. It doesn't set the block entity in the world
func (b *Broadcast) UpdateBlockEntity(x, y, z int32, blockEntity chunk.BlockEntity, dimension string) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	for _, ses := range b.sessions {
		if ses.Player().Dimension() != dimension {
			continue
		}
		ses.UpdateBlockEntity(x, y, z, blockEntity)
	}
}

// returns the number of players in the broadcast
func (b *Broadcast) NumSession() int {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	return len(b.sessions)
}

func (b *Broadcast) DamageEvent(attacker, attacked Session, dimension string, damageType string) {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()

	id := attacked.Player().EntityId()
	sound := EntitySoundEffect(
		"minecraft:entity.player.hurt", false, nil, play.SoundCategoryPlayer, id, 1, 1,
	)

	yawd, _ := attacker.Player().Rotation()
	x, z := YawToXZDelta(yawd)

	vx, vy, vz := attacked.Player().Motion()
	velDeltaX, velDeltaZ := (x+sign(x)*5)/20, (z+sign(z)*5)/20

	vx, vz = vx+velDeltaX, vz+velDeltaZ
	attacked.Player().SetMotion(vx, vy, vz)

	for _, s := range b.sessions {
		if dimension != s.Player().Dimension() {
			continue
		}
		s.DamageEvent(attacker, attacked, damageType)
		s.PlayEntitySound(sound)
		s.(interface {
			Conn() *net.Conn
		}).Conn().WritePacket(&play.SetEntityVelocity{
			EntityId: id,
			X:        int16(vx * 8000),
			Y:        int16(vy * 8000),
			Z:        int16(vz * 8000),
		})
	}
}

func YawToXZDelta(yaw float32) (float64, float64) {
	yawd := float64(yaw)
	radians := yawd * math.Pi / 180.0

	x := -math.Sin(radians)
	z := -math.Cos(radians)

	return x, z
}

func sign(f float64) float64 {
	if f < 0 {
		return -1
	}
	return 1
}

// returns the players to display in status sample
func (b *Broadcast) Sample() []status.StatusSample {
	b.sessions_mu.RLock()
	defer b.sessions_mu.RUnlock()
	samples := make([]status.StatusSample, len(b.sessions))

	var i int
	for _, s := range b.sessions {
		if !s.ClientInformation().AllowServerListing {
			continue
		}
		samples[i] = status.StatusSample{
			Name: s.Username(),
			ID:   s.UUID().String(),
		}
		i++
	}

	samples = samples[:i]

	return samples
}

// creates a sound effect with the provided data. If custom is true, or name wasn't found in the sound registry, the packet will use a custom sound name. This function generates a random seed for this event using level.NewSeed()
func SoundEffect(name string, custom bool, fixedRange *float32, category int32, x, y, z int32, volume, pitch float32) *play.SoundEffect {
	pk := &play.SoundEffect{
		SoundCategory: category,
		X:             x, Y: y, Z: z,
		Volume: volume,
		Pitch:  pitch,
		Seed:   int64(level.NewSeed()),
	}
	if !custom {
		soundId, ok := registry.SoundEvent.Lookup(name)
		if ok {
			pk.SoundId = soundId
			return pk
		}
	}

	pk.SoundId = -1
	pk.SoundName = name
	pk.FixedRange = fixedRange != nil
	if fixedRange != nil {
		pk.Range = *fixedRange
	}

	return pk
}

func EntitySoundEffect(name string, custom bool, fixedRange *float32, category int32, entityId int32, volume, pitch float32) *play.EntitySoundEffect {
	pk := &play.EntitySoundEffect{
		SoundCategory: category,
		EntityId:      entityId,
		Volume:        volume,
		Pitch:         pitch,
		Seed:          int64(level.NewSeed()),
	}
	if !custom {
		soundId, ok := registry.SoundEvent.Lookup(name)
		if ok {
			pk.SoundId = soundId
			return pk
		}
	}

	pk.SoundId = -1
	pk.SoundName = name
	pk.FixedRange = fixedRange != nil
	if fixedRange != nil {
		pk.Range = *fixedRange
	}

	return pk
}
