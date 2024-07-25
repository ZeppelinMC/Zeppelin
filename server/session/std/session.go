package std

import (
	"math"
	nnet "net"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/net/packet/login"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/config"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/player"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/server/world/region"
	"github.com/zeppelinmc/zeppelin/text"
	"github.com/zeppelinmc/zeppelin/util"
)

var _ session.Session = (*StandardSession)(nil)

// StandardSession is a session that uses *net.Conn
type StandardSession struct {
	world     *world.World
	player    *player.Player
	broadcast *session.Broadcast
	config    config.ServerConfig

	conn *net.Conn

	clientName string // constant

	statusProviderProvider func() net.StatusProvider

	hasSessionData atomic.AtomicValue[bool]
	sessionData    atomic.AtomicValue[play.PlayerSession]

	spawned_ents_mu sync.Mutex
	spawnedEntities []int32

	commandManager *command.Manager

	Spawned atomic.AtomicValue[bool]

	registryIndexes map[string][]string

	// the index that should be sent in player chat messages
	ChatIndex atomic.AtomicValue[int32]
	// the previous messages of this player
	prev_msgs_mu     sync.Mutex
	previousMessages []play.PreviousMessage
}

func NewStandardSession(
	conn *net.Conn,
	player *player.Player,
	world *world.World,
	broadcast *session.Broadcast,
	config config.ServerConfig,
	statusProviderProvider func() net.StatusProvider,
	commandManager *command.Manager,
) *StandardSession {
	return &StandardSession{
		conn:                   conn,
		world:                  world,
		player:                 player,
		broadcast:              broadcast,
		config:                 config,
		statusProviderProvider: statusProviderProvider,
		commandManager:         commandManager,

		registryIndexes: make(map[string][]string),
	}
}

func (session *StandardSession) CommandManager() *command.Manager {
	return session.commandManager
}

func (session *StandardSession) WritePacket(pk packet.Packet) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) ReadPacket() (packet.Packet, error) {
	return session.conn.ReadPacket()
}

func (session *StandardSession) SynchronizePosition(x, y, z float64, yaw, pitch float32) error {
	session.player.SetPosition(x, y, z)
	session.player.SetRotation(yaw, pitch)
	return session.conn.WritePacket(&play.SynchronizePlayerPosition{X: x, Y: y, Z: z, Yaw: yaw, Pitch: pitch})
}

func (session *StandardSession) UpdateEntityPosition(pk *play.UpdateEntityPosition) error {
	return session.conn.WritePacket(pk)
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityPositionRotation(pk *play.UpdateEntityPositionAndRotation) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetHeadRotation{EntityId: pk.EntityId, HeadYaw: pk.Yaw})
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityRotation(pk *play.UpdateEntityRotation) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetHeadRotation{EntityId: pk.EntityId, HeadYaw: pk.Yaw})
}

func (session *StandardSession) EntityAnimation(entityId int32, animation byte) error {
	return session.conn.WritePacket(&play.EntityAnimation{EntityId: entityId, Animation: animation})
}

func (session *StandardSession) EntityMetadata(entityId int32, md metadata.Metadata) error {
	return session.conn.WritePacket(&play.SetEntityMetadata{EntityId: entityId, Metadata: md})
}

func (session *StandardSession) Conn() *net.Conn {
	return session.conn
}

func (session *StandardSession) Config() config.ServerConfig {
	return session.config
}

func (session *StandardSession) Broadcast() *session.Broadcast {
	return session.broadcast
}

func (session *StandardSession) Player() *player.Player {
	return session.player
}

func (session *StandardSession) Addr() nnet.Addr {
	return session.conn.RemoteAddr()
}

func (session *StandardSession) ClientName() string {
	return session.clientName
}

func (session *StandardSession) Username() string {
	return session.conn.Username()
}

func (session *StandardSession) UUID() uuid.UUID {
	return session.conn.UUID()
}

func (session *StandardSession) Properties() []login.Property {
	return session.conn.Properties()
}

func (session *StandardSession) SessionData() (play.PlayerSession, bool) {
	return session.sessionData.Get(), session.hasSessionData.Get()
}

func (session *StandardSession) PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) PlayerInfoRemove(uuids ...uuid.UUID) error {
	return session.conn.WritePacket(&play.PlayerInfoRemove{UUIDs: uuids})
}

func (session *StandardSession) Disconnect(reason text.TextComponent) error {
	if session.inConfiguration() {
		return session.conn.WritePacket(&configuration.Disconnect{Reason: reason})
	} else {
		return session.conn.WritePacket(&play.Disconnect{Reason: reason})
	}
}

// finishes configuration
func (session *StandardSession) Configure() error {
	go session.handlePackets()
	if err := session.conn.WritePacket(&configuration.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, session.config.Brand),
	}); err != nil {
		return err
	}
	for _, packet := range configuration.RegistryPackets {
		if err := session.conn.WritePacket(packet); err != nil {
			return err
		}
		session.registryIndexes[packet.RegistryId] = slices.Clone(packet.Indexes)
	}

	if err := session.conn.WritePacket(updateTags); err != nil {
		return err
	}

	if err := session.conn.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}
	return nil
}

func (session *StandardSession) login() error {
	if err := session.conn.WritePacket(&play.Login{
		EntityID:   session.player.EntityId(),
		Dimensions: []string{session.player.Dimension()},

		Hardcore: session.world.Data.Hardcore,

		ViewDistance:       session.config.RenderDistance,
		SimulationDistance: session.config.SimulationDistance,

		HashedSeed: session.world.Data.WorldGenSettings.Seed.HashedSeed(),

		EnableRespawnScreen: true,
		DimensionType:       int32(session.Dimension().Type()),
		DimensionName:       session.player.Dimension(),
		GameMode:            byte(session.player.GameMode()),

		EnforcesSecureChat: session.config.Chat.ChatMode == "secure" || session.config.Chat.DisableWarning,
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.ChangeDifficulty{
		Difficulty: session.world.Data.Difficulty,
		Locked:     session.world.Data.DifficultyLocked,
	}); err != nil {
		return err
	}

	movementSpeed := session.player.Attribute("minecraft:generic.movement_speed")
	if movementSpeed == nil {
		movementSpeed = &entity.Attribute{
			Id:   "minecraft:generic.movement_speed",
			Base: 0.1,
		}
		session.player.SetAttribute("minecraft:generic.movement_speed", 0.1)
	}

	if err := session.conn.WritePacket(
		session.player.Abilities().Encode(float32(movementSpeed.Base)),
	); err != nil {
		return err
	}

	if err := session.conn.WritePacket(session.commandManager.Encode()); err != nil {
		return err
	}

	recipeBook := session.player.RecipeBook()

	if err := session.conn.WritePacket(&play.UpdateRecipeBook{
		Action:                         play.UpdateRecipeBookActionInit,
		CraftingRecipeBookOpen:         recipeBook.IsGuiOpen,
		CraftingRecipeBookFilterActive: recipeBook.IsFilteringCraftable,

		SmeltingRecipeBookOpen:         recipeBook.IsFurnaceGuiOpen,
		SmeltingRecipeBookFilterActive: recipeBook.IsFurnaceFilteringCraftable,

		BlastFurnaceRecipeBookOpen:         recipeBook.IsBlastingFurnaceGuiOpen,
		BlastFurnaceRecipeBookFilterActive: recipeBook.IsBlastingFurnaceFilteringCraftable,

		SmokerRecipeBookOpen:         recipeBook.IsSmokerGuiOpen,
		SmokerRecipeBookFilterActive: recipeBook.IsSmokerFilteringCraftable,

		Array1: recipeBook.ToBeDisplayed,
		Array2: recipeBook.Recipes,
	}); err != nil {
		return err
	}

	x, y, z := session.player.Position()
	yaw, pitch := session.player.Rotation()
	if err := session.SynchronizePosition(x, y, z, yaw, pitch); err != nil {
		return err
	}

	status := session.statusProviderProvider()()

	if err := session.conn.WritePacket(&play.ServerData{
		MOTD: status.Description,
		Icon: status.Favicon,
	}); err != nil {
		return err
	}

	session.broadcast.AddPlayer(session)

	if err := session.WritePacket(&play.SetDefaultSpawnPosition{
		X:     session.world.Data.SpawnX,
		Y:     session.world.Data.SpawnY,
		Z:     session.world.Data.SpawnZ,
		Angle: session.world.Data.SpawnAngle,
	}); err != nil {
		return err
	}

	if err := session.sendSpawnChunks(); err != nil {
		return err
	}

	return nil
}

func (session *StandardSession) IsSpawned(entityId int32) bool {
	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	for _, e := range session.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (session *StandardSession) Dimension() *region.Dimension {
	return session.world.Dimension(session.player.Dimension())
}

/*
The view distance of the client, in chunks

Returns the server's render distance if the client's view distance is bigger or not set
*/
func (session *StandardSession) ViewDistance() int32 {
	plVd := int32(session.player.ClientInformation().ViewDistance)
	if plVd == 0 || plVd > session.config.RenderDistance {
		return session.config.RenderDistance
	}

	return plVd
}

func (session *StandardSession) DespawnEntities(entityIds ...int32) error {
	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	session.spawnedEntities = slices.DeleteFunc(session.spawnedEntities, func(entityId int32) bool {
		for _, e := range entityIds {
			if e == entityId {
				return true
			}
		}
		return false
	})
	return session.conn.WritePacket(&play.RemoveEntities{EntityIDs: entityIds})
}

func (session *StandardSession) bundleDelimiter() error {
	return session.conn.WritePacket(&play.BundleDelimiter{})
}

func (session *StandardSession) SpawnEntity(e entity.Entity) error {
	if err := session.bundleDelimiter(); err != nil {
		return err
	}
	x, y, z := e.Position()
	yaw, pitch := e.Rotation()
	id := e.EntityId()

	if err := session.conn.WritePacket(&play.SpawnEntity{
		EntityId:   id,
		EntityUUID: e.UUID(),
		Type:       e.Type(),
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        util.DegreesToAngle(yaw),
		Pitch:      util.DegreesToAngle(pitch),
		HeadYaw:    util.DegreesToAngle(yaw),
	}); err != nil {
		return err
	}

	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	session.spawnedEntities = append(session.spawnedEntities, id)

	if err := session.conn.WritePacket(&play.SetEntityMetadata{
		EntityId: id,
		Metadata: e.Metadata(),
	}); err != nil {
		return err
	}

	if err := session.bundleDelimiter(); err != nil {
		return err
	}

	return nil
}

func (session *StandardSession) SpawnPlayer(ses session.Session) error {
	return session.SpawnEntity(ses.Player())
}

func (session *StandardSession) sendSpawnChunks() error {
	viewDistance := session.ViewDistance()

	x, _, z := session.player.Position()
	chunkX, chunkZ := int32(math.Floor(x/16)), int32(math.Floor(z/16))

	if err := session.conn.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	var chunks int32

	if err := session.conn.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	for x := chunkX - viewDistance; x <= chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z <= chunkZ+viewDistance; z++ {
			c, err := session.world.Overworld.GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.conn.WritePacket(c.Encode(session.registryIndexes["minecraft:worldgen/biome"])); err != nil {
				return err
			}
			chunks++
		}
	}

	if err := session.conn.WritePacket(&play.ChunkBatchFinished{
		BatchSize: chunks,
	}); err != nil {
		return err
	}

	return nil
}

func (session *StandardSession) UpdateTime(worldAge, dayTime int64) error {
	return session.conn.WritePacket(&play.UpdateTime{WorldAge: worldAge, TimeOfDay: dayTime})
}