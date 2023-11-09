package player

import (
	"math"
	"math/rand"
	"slices"
	"sync"
	"sync/atomic"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/block"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/config"
	"github.com/dynamitemc/dynamite/server/controller"
	"github.com/dynamitemc/dynamite/server/entity"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/inventory"
	"github.com/dynamitemc/dynamite/server/item"
	"github.com/dynamitemc/dynamite/server/lang"
	"github.com/dynamitemc/dynamite/server/permission"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/google/uuid"
)

type Position struct {
	X, Y, Z int32
}

func (p Position) Data() uint64 {
	return ((uint64(p.X) & 0x3FFFFFF) << 38) | ((uint64(p.Z) & 0x3FFFFFF) << 12) | (uint64(p.Y) & 0xFFF)
}

var tags = &packet.UpdateTags{
	Tags: []packet.TagType{
		{
			Type: "minecraft:fluid",
			Tags: []packet.Tag{
				{
					Name:    "minecraft:water",
					Entries: []int32{02, 01},
				},
				{
					Name:    "minecraft:lava",
					Entries: []int32{04, 03},
				},
			},
		},
	},
}

type clientInfo struct {
	Locale               string
	ChatMode             int32
	ChatColors           bool
	DisplayedSkinParts   uint8
	MainHand             int32
	DisableTextFiltering bool
	AllowServerListings  bool
	ViewDistance         int8
}

type Player struct {
	Server any
	config *config.Config
	lang   *lang.Lang
	logger *logger.Logger

	playerController *controller.Controller[uuid.UUID, *Player]
	entityController *controller.Controller[int32, entity.Entity]

	conn *minecraft.Conn

	uuid     uuid.UUID
	entityID int32

	isHardCore *atomic.Bool
	gameMode   byte

	dead           *atomic.Bool
	health         float32
	food           *atomic.Int32
	foodSaturation float32

	data *world.PlayerData

	Inventory            *inventory.Inventory
	previousSelectedSlot item.Item
	selectedSlot         *atomic.Int32

	dimension *world.Dimension

	clientInfo clientInfo

	x, y, z                    float64
	yaw, pitch                 float32
	onGround, operator, flying *atomic.Bool
	highestY                   float64

	spawnedEntities []int32
	loadedChunks    map[[2]int32]struct{}

	sessionID    [16]byte
	publicKey    []byte
	keySignature []byte
	expires      int64

	previousMessages              []packet.PreviousMessage
	acknowledgedMessageSignatures [][]byte
	index                         *atomic.Int32

	newID func() int32

	mu sync.RWMutex
}

func newAtomicInt32(val int32) *atomic.Int32 {
	v := &atomic.Int32{}
	v.Add(val)
	return v
}

func newAtomicBool(val bool) *atomic.Bool {
	v := &atomic.Bool{}
	v.Store(val)
	return v
}

func New(
	uuid uuid.UUID,
	players *controller.Controller[uuid.UUID, *Player],
	entities *controller.Controller[int32, entity.Entity],
	server any,
	config *config.Config,
	lang *lang.Lang,
	logger *logger.Logger,
	entityId int32,
	conn *minecraft.Conn,
	data *world.PlayerData,
	dimension *world.Dimension,
	vd int8,
	newID func() int32,
) *Player {
	pl := &Player{
		Server:           server,
		logger:           logger,
		config:           config,
		lang:             lang,
		playerController: players,
		entityController: entities,
		conn:             conn,
		uuid:             uuid,
		entityID:         entityId,
		isHardCore:       &atomic.Bool{},
		gameMode:         byte(data.PlayerGameType),
		dead:             &atomic.Bool{},
		food:             newAtomicInt32(data.FoodLevel),
		data:             data,
		Inventory:        inventory.From(data.Inventory, data.SelectedItemSlot),
		selectedSlot:     newAtomicInt32(data.SelectedItemSlot),
		dimension:        dimension,
		onGround:         &atomic.Bool{},
		operator:         &atomic.Bool{},
		flying:           newAtomicBool(data.Abilities.Flying&int8(enum.PlayerAbilityFlying) != 0),
		index:            &atomic.Int32{},
		x:                data.Pos[0],
		y:                data.Pos[1],
		z:                data.Pos[2],
		yaw:              data.Rotation[0],
		pitch:            data.Rotation[1],
		health:           data.Health,
		foodSaturation:   data.FoodSaturationLevel,
		newID:            newID,
	}
	pl.clientInfo.ViewDistance = vd

	fl := true
	if data.Abilities.Flying == 0 {
		fl = false
	}
	pl.flying = newAtomicBool(fl)

	return pl
}

func (p *Player) UUID() uuid.UUID {
	return p.uuid
}

func (p *Player) Name() string {
	return p.conn.Name()
}

func (p *Player) EntityID() int32 {
	return p.entityID
}

func (p *Player) SendPacket(pk packet.Packet) error {
	return p.conn.SendPacket(pk)
}

func (p *Player) ReadPacket() (packet.Packet, error) {
	return p.conn.ReadPacket()
}

func (p *Player) Save() {
	o := int8(0)
	if p.onGround.Load() {
		o = 1
	}
	fl := int8(0)
	if p.flying.Load() {
		fl = 1
	}

	p.mu.Lock()
	p.data.Pos[0], p.data.Pos[1], p.data.Pos[2], p.data.Rotation[0], p.data.Rotation[1], p.data.OnGround = p.x, p.y, p.z, p.yaw, p.pitch, o
	p.data.PlayerGameType = int32(p.gameMode)
	p.data.Inventory = p.Inventory.Data()
	p.data.Abilities.Flying = fl
	p.data.Dimension = p.dimension.Type()
	p.data.SelectedItemSlot = p.selectedSlot.Load()
	p.data.Health = p.health
	p.data.FoodSaturationLevel = p.foodSaturation
	p.data.FoodLevel = p.food.Load()
	p.mu.Unlock()

	p.data.Save()
}

func (p *Player) Respawn(d *world.Dimension) {
	p.SetDead(false)
	p.SendPacket(&packet.Respawn{
		GameMode:         p.GameMode(),
		PreviousGameMode: -1,
		DimensionType:    d.Type(),
		DimensionName:    d.Type(),
		HashedSeed:       d.Seed(),
	})
	p.SetDimension(d)

	var x1, y1, z1 int32
	var a float32

	switch d.Type() {
	case "minecraft:overworld":
		x1, y1, z1, a = d.World().Spawn()
	case "minecraft:the_nether":
		x, y, z := p.Position()
		x1, y1, z1 = int32(x)/8, int32(y)/8, int32(z)/8
	}

	clear(p.spawnedEntities)

	yaw, pitch := p.Rotation()

	if b, _ := world.GameRule(d.World().Gamerules()["keepInventory"]).Bool(); !b {
		p.Inventory.Clear()
	}

	p.SendPacket(&packet.SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.Inventory.Packet(),
	})

	chunkX, chunkZ := math.Floor(float64(x1)/16), math.Floor(float64(z1)/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)
	p.SendChunks(d)

	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)

	p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: Position{x1, y1, z1}.Data(),
		Angle:    a,
	})
}

func (p *Player) Login(d *world.Dimension) {
	p.SendPacket(&packet.JoinGame{
		EntityID:           p.entityID,
		IsHardcore:         p.IsHardcore(),
		GameMode:           p.GameMode(),
		PreviousGameMode:   -1,
		DimensionNames:     []string{"minecraft:overworld", "minecraft:the_nether", "minecraft:the_end"},
		DimensionType:      d.Type(),
		DimensionName:      d.Type(),
		HashedSeed:         d.Seed(),
		MaxPlayers:         0,
		ViewDistance:       int32(p.clientInfo.ViewDistance),
		SimulationDistance: int32(p.clientInfo.ViewDistance), //todo fix this
	})
	p.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite"),
	})
	x1, y1, z1 := p.Position()
	yaw, pitch := p.Rotation()

	chunkX, chunkZ := math.Floor(x1/16), math.Floor(z1/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})

	p.SetPosition(x1, y1, z1, yaw, pitch, false)
	p.SendChunks(d)

	abs := p.SavedAbilities()
	abps := &packet.PlayerAbilities{FlyingSpeed: abs.FlySpeed, FieldOfViewModifier: 0.1}
	if abs.Flying != 0 {
		abps.Flags |= enum.PlayerAbilityFlying | enum.PlayerAbilityAllowFlying
	}
	if abps.Flags != 0 {
		p.SendPacket(abps)
	}
	p.SendPacket(tags)

	if p.Operator() {
		p.SendPacket(&packet.EntityEvent{
			EntityID: p.entityID,
			Status:   28,
		})
	}

	x, y, z, a := d.World().Spawn()
	p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: Position{x, y, z}.Data(),
		Angle:    a,
	})

	p.Teleport(x1, y1, z1, yaw, pitch)

	/*if p.Server.Config.ResourcePack.Enable {
		p.SendPacket(&packet.ResourcePack{
			URL:    p.Server.Config.ResourcePack.URL,
			Hash:   p.Server.Config.ResourcePack.Hash,
			Forced: p.Server.Config.ResourcePack.Force,
			//Prompt: p.Server.Config.Messages.ResourcePackPrompt,
		})
	}*/
}

func (p *Player) SendMessage(message chat.Message) error {
	return p.SendPacket(&packet.SystemChatMessage{Message: message})
}

func (p *Player) Damage(health float32) {
	p.SendPacket(&packet.EntitySoundEffect{
		Category: enum.SoundCategoryAmbient,
		SoundID:  519,
		EntityID: p.entityID,
		Seed:     world.RandomSeed(),
		Volume:   1,
		Pitch:    1,
	})
	p.SetHealth(p.Health() - health)
}

func (p *Player) Kill(message string) {
	p.SetDead(true)
	if f, _ := world.GameRule(p.Dimension().World().Gamerules()["doImmediateRespawn"]).Bool(); !f {
		p.SendPacket(&packet.GameEvent{
			Event: uint8(enum.GameEventEnableRespawnScreen),
			Value: 0,
		})
	}

	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		if !pl.IsSpawned(p.entityID) {
			return true
		}
		pl.SendPacket(&packet.DamageEvent{
			EntityID:     p.entityID,
			SourceTypeID: int32(enum.DamageTypeGenericKill),
		})
		pl.SendPacket(&packet.EntityEvent{
			EntityID: p.entityID,
			Status:   3,
		})
		return true
	})

	p.SendPacket(&packet.DamageEvent{
		EntityID:     p.entityID,
		SourceTypeID: int32(enum.DamageTypeGenericKill),
	})
	p.Despawn()
	p.SendPacket(&packet.CombatDeath{
		Message:  message,
		PlayerID: p.entityID,
	})
	p.SetHealth(20)
	p.SetFoodLevel(20)
	p.SetFoodSaturationLevel(5)
}

func (p *Player) Push(x, y, z float64) {
	yaw, pitch := p.Rotation()
	p.SetPosition(x, y, z, yaw, pitch, p.OnGround())
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: p.newID(),
	})
	//p.BroadcastMovement(0x14, x, y, z, yaw, pitch, p.Player.OnGround(), false)
}

func (p *Player) Teleport(x, y, z float64, yaw, pitch float32) {
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: p.newID(),
	})
	p.BroadcastMovement(0, x, y, z, yaw, pitch, p.OnGround(), true)
}

func (p *Player) SendCommands(g *commands.Graph) {
	graph := commands.Graph{
		Commands: make([]*commands.Command, len(g.Commands)),
	}
	copy(graph.Commands, g.Commands)
	for i, command := range graph.Commands {
		if command == nil {
			continue
		}
		if !p.HasPermissions(command.RequiredPermissions) {
			graph.Commands[i] = nil
		}
	}
	p.SendPacket(graph.Data())
}

func (p *Player) Keepalive() {
	id := rand.Int63() * 100
	p.SendPacket(&packet.KeepAliveClient{PayloadID: id})
}

func (p *Player) Disconnect(reason chat.Message) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.SendPacket(pk)
	p.conn.Close(nil)
}

func (p *Player) IsChunkLoaded(x, z int32) bool {
	_, ok := p.loadedChunks[[2]int32{x, z}]
	return ok
}

func (p *Player) SendChunks(dimension *world.Dimension) {
	x1, _, z1 := p.Position()
	p.mu.Lock()
	defer p.mu.Unlock()
	max := int32(p.clientInfo.ViewDistance)
	if p.loadedChunks == nil {
		p.loadedChunks = make(map[[2]int32]struct{})
	}

	chunkX := int32(math.Floor(x1 / 16))
	chunkZ := int32(math.Floor(z1 / 16))

	for x := chunkX - max; x <= chunkX+max; x++ {
		for z := chunkZ - max; z <= chunkZ+max; z++ {
			if p.IsChunkLoaded(x, z) {
				continue
			}
			c, err := dimension.Chunk(x, z)
			if err != nil {
				continue
			}
			p.loadedChunks[[2]int32{x, z}] = struct{}{}
			p.SendPacket(c.Data())

			for _, en := range c.Entities {
				u, _ := world.IntUUIDToByteUUID(en.UUID)

				var e entity.Entity

				if f := findEntityByUUID(p.entityController, p.playerController, u); f != nil {
					if d, ok := f.(entity.Entity); ok {
						e = d
					}
				} else {
					e = entity.CreateEntity(p.entityController, p.newID(), en, dimension)
				}

				t, ok := registry.GetEntity(e.Type())
				if !ok {
					continue
				}

				x, y, z := e.Position()
				yaw, pitch := e.Rotation()
				p.SendPacket(&packet.SpawnEntity{
					EntityID: e.EntityID(),
					UUID:     u,
					X:        x,
					Y:        y,
					Z:        z,
					Pitch:    DegreesToAngle(yaw),
					Yaw:      DegreesToAngle(pitch),
					Type:     t.ProtocolID,
				})
				p.spawnedEntities = append(p.spawnedEntities, e.EntityID())
			}
		}
	}
}

func (p *Player) ChunkPosition() (int32, int32) {
	x, _, z := p.Position()
	return int32(x) / 16, int32(z) / 16
}

func (p *Player) GetPrefixSuffix() (prefix string, suffix string) {
	group := permission.GetGroup(permission.GetPlayer(p.UUID().String()).Group)
	return group.Prefix, group.Suffix
}

func (p *Player) HandleCenterChunk(x1, z1, x2, z2 float64) {
	oldChunkX := int(math.Floor(x1 / 16))
	oldChunkZ := int(math.Floor(z1 / 16))

	newChunkX := int(math.Floor(x2 / 16))
	newChunkZ := int(math.Floor(z2 / 16))

	if newChunkX != oldChunkX || newChunkZ != oldChunkZ {
		p.SendPacket(&packet.SetCenterChunk{
			ChunkX: int32(newChunkX),
			ChunkZ: int32(newChunkZ),
		})
	}
}

func (p *Player) IsSpawned(entityId int32) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, e := range p.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (p *Player) SpawnPlayer(pl *Player) {
	entityId := pl.entityID
	x, y, z := pl.Position()
	ya, pi := pl.Rotation()
	yaw, pitch := DegreesToAngle(ya), DegreesToAngle(pi)

	p.SendPacket(&packet.SpawnPlayer{
		EntityID:   entityId,
		PlayerUUID: pl.conn.UUID(),
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
	})
	p.SendPacket(&packet.EntityHeadRotation{
		EntityID: entityId,
		HeadYaw:  yaw,
	})

	p.mu.Lock()
	p.spawnedEntities = append(p.spawnedEntities, entityId)
	p.mu.Unlock()

	pl.sendEquipment(p)
	p.sendEquipment(pl)
}

func (p *Player) DespawnEntity(entityId int32) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.SendPacket(&packet.DestroyEntities{
		EntityIds: []int32{entityId},
	})
	index := -1
	for i, e := range p.spawnedEntities {
		if e == entityId {
			index = i
		}
	}
	if index != -1 {
		p.spawnedEntities = slices.Delete(p.spawnedEntities, index, index+1)
	}
}

func (p *Player) IntitializeData() {
	p.SendPacket(&packet.SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.Inventory.Packet(),
	})
	p.SendPacket(&packet.SetHeldItem{Slot: int8(p.SelectedSlot())})
	p.SendPacket(&packet.SetHealth{
		Health:         p.Health(),
		FoodSaturation: p.FoodSaturationLevel(),
		Food:           p.FoodLevel(),
	})
}

func (p *Player) ClearItem(slot int8) {
	p.SendPacket(&packet.SetContainerSlot{
		WindowID: 0,
		StateID:  1,
		Slot:     int16(inventory.DataSlotToNetworkSlot(slot)),
	})
	p.Inventory.DeleteSlot(slot)
}

func (p *Player) SetSlot(slot int8, data item.Item) {
	s, _ := item.ItemToPacketSlot(data)
	p.SendPacket(&packet.SetContainerSlot{
		WindowID: 0,
		StateID:  1,
		Slot:     int16(inventory.DataSlotToNetworkSlot(slot)),
		Data:     s,
	})
	p.Inventory.SetSlot(slot, data)
}

func (p *Player) DropSlot() {
	item := p.PreviousSelectedSlot()
	s, _ := item.ToPacketSlot()
	x, y, z := p.Position()

	id := p.newID()
	u := uuid.New()

	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		if !pl.InView(p.Position()) {
			return true
		}
		pl.SendPacket(&packet.SpawnEntity{
			EntityID: id,
			UUID:     u,
			Type:     54,
			X:        x,
			Y:        y,
			Z:        z,
		})
		pl.SendPacket(&packet.SetEntityMetadata{
			EntityID: id,
			Slot:     &s,
		})
		return true
	})
}
func (p *Player) SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch) {
	p.SendPacket(&packet.CommandSuggestionsResponse{
		TransactionId: id,
		Start:         start,
		Length:        length,
		Matches:       matches,
	})
}

func (p *Player) OnBlock() block.Block {
	x, y, z := p.Position()
	return p.Dimension().Block(int64(x), int64(y-1), int64(z))
}

func (p *Player) TeleportToEntity(uuid [16]byte) {
	e := findEntityByUUID(p.entityController, p.playerController, uuid)
	if e == nil {
		return
	}
	if pl, ok := e.(*Player); ok {
		x, y, z := pl.Position()
		yaw, pitch := pl.Rotation()
		p.Teleport(x, y, z, yaw, pitch)
	} else {
		e := e.(entity.Entity)
		x, y, z := e.Position()
		yaw, pitch := e.Rotation()
		p.Teleport(x, y, z, yaw, pitch)
	}
}

func (p *Player) IP() string {
	return p.conn.RemoteAddr().String()
}

func (s *Player) HasPermissions(perms []string) bool {
	if s.Operator() {
		return true
	}

	return permission.HasPermissions(s.Name(), perms)
}

func (p *Player) InView(x2, y2, z2 float64) bool {
	x1, y1, z1 := p.Position()
	distance := math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2))

	return float64(p.clientInfo.ViewDistance)*16 > distance
}

func (p *Player) SpawnEntity(pk *packet.SpawnEntity) {
	p.mu.Lock()
	p.spawnedEntities = append(p.spawnedEntities, pk.EntityID)
	p.mu.Unlock()
	p.SendPacket(pk)
}

func (p *Player) IsMessageCached(s [256]byte) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, sig := range p.acknowledgedMessageSignatures {
		if [256]byte(sig) == s {
			return true
		}
	}
	return false
}

func (p *Player) CacheMessage(s []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.acknowledgedMessageSignatures = append(p.acknowledgedMessageSignatures, s)
}

func (p *Player) Index() int32 {
	return p.index.Load()
}

func (p *Player) PreviousMessages() []packet.PreviousMessage {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.previousMessages
}

func (p *Player) AddMessage(sig []byte) {
	p.mu.Lock()
	if len(p.previousMessages) != 20 {
		p.previousMessages = append([]packet.PreviousMessage{
			{
				MessageID: p.index.Load(),
				Signature: sig,
			},
		}, p.previousMessages...)
	}
	p.mu.Unlock()
	p.index.Add(1)
}
