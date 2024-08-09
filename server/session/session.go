package session

import (
	"net"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/net/packet/login"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/properties"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/player"

	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/dimension/window"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/text"
)

type Session interface {
	// Username of the session
	Username() string
	// UUID of the session
	UUID() uuid.UUID
	// Properties (typically only textures)
	Properties() []login.Property

	// The player this session holds
	Player() *player.Player
	// The client name this session reports in minecraft:brand (vanilla)
	ClientName() string
	// the client settings of this client
	ClientInformation() configuration.ClientInformation
	// The address of this connection
	Addr() net.Addr

	// the broadcaster of the session
	Broadcast() *Broadcast

	// the server properties used by this session
	Config() properties.ServerProperties

	// Disconnects the session from the server
	Disconnect(reason text.TextComponent) error
	// sends a player chat message packet to the session
	PlayerChatMessage(pk play.ChatMessage, sender Session, chatType string, index int32, prevMsgs []play.PreviousMessage) error
	// sends a player info update packet to the session
	PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error
	// sends a player info remove packet to the session
	PlayerInfoRemove(uuids ...uuid.UUID) error

	// sends a disguised chat message
	DisguisedChatMessage(content text.TextComponent, sender Session, chatType string) error

	UpdateEntityPosition(entity entity.Entity, pk *play.UpdateEntityPosition) error
	UpdateEntityPositionRotation(entity entity.Entity, pk *play.UpdateEntityPositionAndRotation) error
	UpdateEntityRotation(entity entity.Entity, pk *play.UpdateEntityRotation) error

	// whether the entity is spawned for this session or not
	IsSpawned(entityId int32) bool
	// despawns the entities for this session
	DespawnEntities(entityIds ...int32) error
	// spawns the entity for this session
	SpawnEntity(entity.Entity) error
	// spawns a player
	SpawnPlayer(Session) error

	// whether this session should be listed in the playerlist
	Listed() bool

	// returns the latency of the client in milliseconds
	Latency() int64

	// sends entity animation
	EntityAnimation(entityId int32, animation byte) error
	// sends entity metadata
	EntityMetadata(entityId int32, md metadata.Metadata) error

	// teleports the player to specified location with specified rotation
	SynchronizePosition(x, y, z float64, yaw, pitch float32) error
	// sends a system (unsigned) chat message to the client
	SystemMessage(msg text.TextComponent) error

	// Returns the session data for this session, and if it has any
	SessionData() (data play.PlayerSession, ok bool)

	// updates the time for the client
	UpdateTime(worldAge, dayTime int64) error

	// sets the gamemode for the client
	SetGameMode(gm level.GameMode) error

	// sends a block action to the session
	BlockAction(*play.BlockAction) error

	// opens the window for the client and bumps up the window viewer number
	OpenWindow(w *window.Window) error

	// plays the sound for the client
	PlaySound(*play.SoundEffect) error

	// plays the sound for the client
	PlayEntitySound(*play.EntitySoundEffect) error

	// updates the block for the player
	UpdateBlock(pos pos.BlockPosition, b section.Block) error

	// updates the block entity for the player
	UpdateBlockEntity(pos pos.BlockPosition, be chunk.BlockEntity) error

	// sent damage event to the client
	DamageEvent(attacker, attacked Session, damageType string) error

	// deletes a signed message for the client
	DeleteMessage(id int32, sig [256]byte) error

	SetTickState(tps float32, frozen bool) error

	// the textures of this client
	Textures() (login.Textures, error)
}
