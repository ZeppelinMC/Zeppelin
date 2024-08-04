package std

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/server/world/block"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
	"github.com/zeppelinmc/zeppelin/server/world/dimension/window"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/server/world/region"
	"github.com/zeppelinmc/zeppelin/text"
	"github.com/zeppelinmc/zeppelin/util"
)

var _ session.Session = (*StandardSession)(nil)

// StandardSession is a session that uses *net.Conn
type StandardSession struct {
	World     *world.World
	player    *player.Player
	broadcast *session.Broadcast
	config    config.ServerConfig

	conn *net.Conn

	clientName string // constant
	ClientInfo atomic.AtomicValue[configuration.ClientInformation]

	statusProviderProvider func() net.StatusProvider

	hasSessionData atomic.AtomicValue[bool]
	sessionData    atomic.AtomicValue[play.PlayerSession]

	spawned_ents_mu sync.Mutex
	spawnedEntities []int32

	commandManager *command.Manager

	registryIndexes map[string][]string

	// the index that should be sent in player chat messages
	ChatIndex atomic.AtomicValue[int32]
	// the previous messages of this player
	prev_msgs_mu     sync.Mutex
	previousMessages []play.PreviousMessage

	inBundle atomic.AtomicValue[bool]

	AwaitingTeleportAcknowledgement   atomic.AtomicValue[bool]
	awaitingChunkBatchAcknowledgement atomic.AtomicValue[bool]

	// the time in milliseconds that the keep alive packet was sent to the server from the client
	sbLastKeepalive atomic.AtomicValue[int64]
	// the time in milliseconds that the keep alive packet was sent to the client from the server
	cbLastKeepAlive atomic.AtomicValue[int64]

	load_ch_mu   sync.RWMutex
	loadedChunks map[uint64]bool

	listed atomic.AtomicValue[bool]

	// the window id the client is viewing currently, 0 if none (inventory)
	WindowView atomic.AtomicValue[int32]
}

func (s *StandardSession) Latency() int64 {
	return s.sbLastKeepalive.Get() - s.cbLastKeepAlive.Get()
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
		World:                  world,
		player:                 player,
		broadcast:              broadcast,
		config:                 config,
		statusProviderProvider: statusProviderProvider,
		commandManager:         commandManager,
		loadedChunks:           make(map[uint64]bool),

		listed: atomic.Value(true),

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

func (session *StandardSession) Listed() bool {
	return session.listed.Get()
}

func (session *StandardSession) SetListed(v bool) {
	session.listed.Set(v)
}

func (session *StandardSession) SynchronizePosition(x, y, z float64, yaw, pitch float32) error {
	session.player.SetPosition(x, y, z)
	session.player.SetRotation(yaw, pitch)
	session.AwaitingTeleportAcknowledgement.Set(true)
	return session.conn.WritePacket(&play.SynchronizePlayerPosition{X: x, Y: y, Z: z, Yaw: yaw, Pitch: pitch})
}

func (session *StandardSession) UpdateEntityPosition(entity entity.Entity, pk *play.UpdateEntityPosition) error {
	return session.conn.WritePacket(pk)
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityPositionRotation(entity entity.Entity, pk *play.UpdateEntityPositionAndRotation) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetHeadRotation{EntityId: pk.EntityId, HeadYaw: pk.Yaw})
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityRotation(entity entity.Entity, pk *play.UpdateEntityRotation) error {
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

func (session *StandardSession) ClientInformation() configuration.ClientInformation {
	return session.ClientInfo.Get()
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
		session.conn.WritePacket(&configuration.Disconnect{Reason: reason})
	} else {
		session.conn.WritePacket(&play.Disconnect{Reason: reason})
	}
	return session.conn.Close()
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

		Hardcore: session.World.Data.Hardcore,

		ViewDistance:       session.config.RenderDistance,
		SimulationDistance: session.config.SimulationDistance,

		HashedSeed: session.World.Data.WorldGenSettings.Seed.HashedSeed(),

		EnableRespawnScreen: true,
		DimensionType:       int32(slices.Index(session.registryIndexes["minecraft:dimension_type"], session.Dimension().Type())),
		DimensionName:       session.player.Dimension(),
		GameMode:            byte(session.player.GameMode()),

		EnforcesSecureChat: session.config.Chat.ChatMode == "secure" || session.config.Chat.DisableWarning,
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.ChangeDifficulty{
		Difficulty: session.World.Data.Difficulty,
		Locked:     session.World.Data.DifficultyLocked,
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

	status := session.statusProviderProvider()()

	if err := session.conn.WritePacket(&play.ServerData{
		MOTD: status.Description,
		Icon: status.Favicon,
	}); err != nil {
		return err
	}

	session.broadcast.AddPlayer(session)

	if err := session.SendInventory(); err != nil {
		return err
	}
	if err := session.conn.WritePacket(&play.SetHeldItemClientbound{Slot: int8(session.player.SelectedItemSlot())}); err != nil {
		return err
	}

	if err := session.WritePacket(&play.SetDefaultSpawnPosition{
		X:     session.World.Data.SpawnX,
		Y:     session.World.Data.SpawnY,
		Z:     session.World.Data.SpawnZ,
		Angle: session.World.Data.SpawnAngle,
	}); err != nil {
		return err
	}

	if err := session.sendSpawnChunks(); err != nil {
		return err
	}

	if err := session.SynchronizePosition(x, y, z, yaw, pitch); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	session.broadcast.SpawnPlayer(session)

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

func (session *StandardSession) Dimension() *dimension.Dimension {
	return session.World.Dimension(session.player.Dimension())
}

/*
The view distance of the client, in chunks

Returns the server's render distance if the client's view distance is bigger or not set
*/
func (session *StandardSession) ViewDistance() int32 {
	plVd := int32(session.ClientInformation().ViewDistance)
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

func (session *StandardSession) bundleStart() error {
	if session.inBundle.Get() {
		return nil
	}
	session.inBundle.Set(true)
	return session.conn.WritePacket(&play.BundleDelimiter{})
}

func (session *StandardSession) bundleStop() error {
	if !session.inBundle.Get() {
		return nil
	}
	session.inBundle.Set(false)
	return session.conn.WritePacket(&play.BundleDelimiter{})
}

func (session *StandardSession) SpawnEntity(e entity.Entity) error {
	if err := session.bundleStart(); err != nil {
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

	if err := session.bundleStop(); err != nil {
		return err
	}

	return nil
}

func (session *StandardSession) UpdateBlock(pos pos.BlockPosition, b block.Block) error {
	state, ok := section.BlockStateId(b)
	if !ok {
		return fmt.Errorf("invalid block")
	}
	x, y, z := pos.X(), pos.Y(), pos.Z()
	return session.WritePacket(&play.BlockUpdate{
		X: x, Y: y, Z: z,
		BlockId: state,
	})
}

func (session *StandardSession) UpdateBlockEntity(pos pos.BlockPosition, be chunk.BlockEntity) error {
	id, ok := registry.BlockEntityType.Lookup(be.Id)
	if !ok {
		return fmt.Errorf("invalid block entity")
	}
	x, y, z := pos.X(), pos.Y(), pos.Z()
	return session.WritePacket(&play.BlockEntityData{
		X: x, Y: y, Z: z,
		Type: id,
		Data: be,
	})
}

func (session *StandardSession) SpawnPlayer(ses session.Session) error {
	return session.SpawnEntity(ses.Player())
}

func (session *StandardSession) SetGameMode(gm level.GameMode) error {
	session.player.SetGameMode(gm)
	return session.WritePacket(&play.GameEvent{
		Event: play.GameEventChangeGamemode,
		Value: float32(gm),
	})
}

func (session *StandardSession) sendSpawnChunks() error {
	viewDistance := session.ViewDistance()

	x, _, z := session.player.Position()
	chunkX, chunkZ := int32(x)>>4, int32(z)>>4

	if err := session.conn.WritePacket(&play.SetCenterChunk{ChunkX: chunkX, ChunkZ: chunkZ}); err != nil {
		return err
	}

	var chunks int32

	if err := session.conn.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	for x := chunkX - viewDistance; x < chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
			c, err := session.Dimension().GetChunk(x, z)
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

	session.awaitingChunkBatchAcknowledgement.Set(true)

	return nil
}

func (session *StandardSession) DamageEvent(attacker, attacked session.Session, damageType string) error {
	causeType := slices.Index(session.registryIndexes["minecraft:damage_type"], damageType)
	if causeType == -1 {
		return fmt.Errorf("unknown damage type")
	}
	de := &play.DamageEvent{
		SourceTypeId: int32(causeType),
	}
	if attacker != nil {
		id := attacker.Player().EntityId()
		de.SourceCauseId = id
		de.SourceDirectId = id
	}
	if err := session.conn.WritePacket(de); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.HurtAnimation{
		EntityId: attacked.Player().EntityId(),
	})
}

func (session *StandardSession) SendChunkRadius(chunkX, chunkZ int32) error {
	viewDistance := session.ViewDistance()

	session.load_ch_mu.Lock()
	defer session.load_ch_mu.Unlock()

	for x := chunkX - viewDistance; x < chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
			if session.loadedChunks[region.ChunkHash(x, z)] {
				continue
			}
			session.loadedChunks[region.ChunkHash(x, z)] = true
			c, err := session.Dimension().GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.conn.WritePacket(c.Encode(session.registryIndexes["minecraft:worldgen/biome"])); err != nil {
				return err
			}
		}
	}

	return nil
}

func (session *StandardSession) UpdateTime(worldAge, dayTime int64) error {
	return session.conn.WritePacket(&play.UpdateTime{WorldAge: worldAge, TimeOfDay: dayTime})
}

func (session *StandardSession) DeleteMessage(id int32, sig [256]byte) error {
	return session.conn.WritePacket(&play.DeleteMessage{MessageId: id, Signature: sig})
}

func (session *StandardSession) SendInventory() error {
	return session.conn.WritePacket(&play.SetContainerContent{
		StateId: 1,
		Slots:   session.player.Inventory().EncodeResize(46),
	})
}

func (session *StandardSession) setContainerContent(container window.Window) error {
	return session.conn.WritePacket(&play.SetContainerContent{
		StateId:  1,
		WindowID: byte(container.Id),
		Slots:    container.Items.Encode(),
	})
}

func (session *StandardSession) OpenWindow(w *window.Window) error {
	if curw := session.WindowView.Get(); curw != 0 {
		return fmt.Errorf("window %d already open for client", curw)
	}
	session.WindowView.Set(w.Id)

	err := session.conn.WritePacket(&play.OpenScreen{
		WindowId:    w.Id,
		WindowType:  registry.Menu.Get(w.Type),
		WindowTitle: w.Title,
	})
	if err != nil {
		return err
	}

	w.Viewers++

	return session.setContainerContent(*w)
}

func (session *StandardSession) Textures() (login.Textures, error) {
	var textures login.Textures
	properties := session.conn.Properties()
	if len(properties) == 0 {
		return textures, fmt.Errorf("client has no textures")
	}
	property := properties[0].Value
	data, err := base64.StdEncoding.DecodeString(property)
	if err != nil {
		return textures, err
	}
	err = json.Unmarshal(data, &textures)
	return textures, err
}

func (session *StandardSession) BlockAction(pk *play.BlockAction) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) PlaySound(pk *play.SoundEffect) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) PlayEntitySound(pk *play.EntitySoundEffect) error {
	return session.conn.WritePacket(pk)
}
