package server

import (
	"math"
	"math/rand"
	"slices"
	"strings"
	"sync"

	"github.com/aimjel/minecraft"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/registry"
	"github.com/dynamitemc/dynamite/server/world"
)

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

type PlayerController struct {
	mu     sync.RWMutex
	player *player.Player
	conn   *minecraft.Conn
	Server *Server

	spawnedEntities []int32
	loadedChunks    map[[2]int32]struct{}

	entityID int32

	UUID string

	clientInfo struct {
		Locale               string
		ViewDistance         int8
		ChatMode             int32
		ChatColors           bool
		DisplayedSkinParts   uint8
		MainHand             int32
		DisableTextFiltering bool
		AllowServerListings  bool
	}
}

func (p *PlayerController) Name() string {
	return p.conn.Name()
}

func (p *PlayerController) SendPacket(pk packet.Packet) error {
	return p.conn.SendPacket(pk)
}

func (p *PlayerController) SetClientSettings(pk *packet.ClientSettings) {
	p.clientInfo.Locale = pk.Locale
	//don't set view distance but server controls it
	p.clientInfo.ChatMode = pk.ChatMode
	p.clientInfo.ChatColors = pk.ChatColors
	p.clientInfo.DisplayedSkinParts = pk.DisplayedSkinParts
	p.clientInfo.MainHand = pk.MainHand
	p.clientInfo.DisableTextFiltering = pk.DisableTextFiltering
	p.clientInfo.AllowServerListings = pk.AllowServerListings
}

func (p *PlayerController) Respawn(dim string) {
	d := p.Server.GetDimension(dim)
	p.player.SetDead(false)
	p.SendPacket(&packet.Respawn{
		GameMode:         p.player.GameMode(),
		PreviousGameMode: -1,
		DimensionType:    d.Type(),
		DimensionName:    d.Type(),
		HashedSeed:       d.Seed(),
	})
	p.player.SetDimension(d.Type())

	var x1, y1, z1 int32
	var a float32

	switch d.Type() {
	case "minecraft:overworld":
		x1, y1, z1, a = p.Server.World.Spawn()
	case "minecraft:the_nether":
		x, y, z := p.Position()
		x1, y1, z1 = int32(x)/8, int32(y)/8, int32(z)/8
	}

	yaw, pitch := p.player.Rotation()

	if b, _ := world.GameRule(p.Server.World.Gamerules()["keepInventory"]).Bool(); b {
		p.InitializeInventory()
	}

	chunkX, chunkZ := math.Floor(float64(x1)/16), math.Floor(float64(z1)/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)
	p.SendSpawnChunks(d)

	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)

	p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: ((uint64(x1) & 0x3FFFFFF) << 38) | ((uint64(z1) & 0x3FFFFFF) << 12) | (uint64(y1) & 0xFFF),
		Angle:    a,
	})
}

func (p *PlayerController) Login(dim string) error {
	d := p.Server.GetDimension(dim)
	if err := p.SendPacket(&packet.JoinGame{
		EntityID:           p.entityID,
		IsHardcore:         p.player.IsHardcore(),
		GameMode:           p.player.GameMode(),
		PreviousGameMode:   -1,
		DimensionNames:     []string{"minecraft:overworld", "minecraft:the_nether", "minecraft:the_end"},
		DimensionType:      d.Type(),
		DimensionName:      d.Type(),
		HashedSeed:         d.Seed(),
		MaxPlayers:         0,
		ViewDistance:       int32(p.clientInfo.ViewDistance),
		SimulationDistance: int32(p.clientInfo.ViewDistance), //todo fix this
	}); err != nil {
		return err
	}
	p.player.SetDimension(d.Type())
	p.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite"),
	})

	x1, y1, z1 := p.player.Position()
	yaw, pitch := p.player.Rotation()

	chunkX, chunkZ := math.Floor(x1/16), math.Floor(z1/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(x1, y1, z1, yaw, pitch)
	p.SendSpawnChunks(d)

	abs := p.player.SavedAbilities()
	abps := &packet.PlayerAbilities{FlyingSpeed: abs.FlySpeed, FieldOfViewModifier: 0.1}
	if abs.Flying != 0 {
		abps.Flags |= 0x06
	}
	if p.player.GameMode() == 1 {
		if abps.Flags == 0 {
			abps.Flags |= 0x04
		}
		abps.Flags |= 0x08
	}

	if abps.Flags != 0 {
		p.SendPacket(abps)
	}
	p.SendPacket(tags)

	if p.player.Operator() {
		p.SendPacket(&packet.EntityEvent{
			EntityID: p.entityID,
			Status:   28,
		})
	}

	p.Teleport(x1, y1, z1, yaw, pitch)

	x, y, z, a := p.Server.World.Spawn()

	v := p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: ((uint64(x) & 0x3FFFFFF) << 38) | ((uint64(z) & 0x3FFFFFF) << 12) | (uint64(y) & 0xFFF),
		Angle:    a,
	})

	if p.Server.Config.ResourcePack.Enable {
		return p.SendPacket(&packet.ResourcePack{
			URL:    p.Server.Config.ResourcePack.URL,
			Hash:   p.Server.Config.ResourcePack.Hash,
			Forced: p.Server.Config.ResourcePack.Force,
			Prompt: p.Server.Config.Messages.ResourcePackPrompt,
		})
	}
	return v
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) SetHealth(health float32) {
	p.player.SetHealth(health)
	food, saturation := p.player.FoodLevel(), p.player.FoodSaturationLevel()
	p.SendPacket(&packet.SetHealth{
		Health:         health,
		Food:           food,
		FoodSaturation: saturation,
	})
	if health == 0 {
		p.Kill("died :skull:")
	}
}

func (p *PlayerController) Kill(message string) {
	p.player.SetDead(true)
	p.BroadcastHealth()
	if f, _ := world.GameRule(p.Server.World.Gamerules()["doImmediateRespawn"]).Bool(); !f {
		p.SendPacket(&packet.GameEvent{
			Event: 11,
			Value: 0,
		})
	}
	p.BroadcastPacketAll(&packet.DamageEvent{
		EntityID:     p.entityID,
		SourceTypeID: 0,
	})
	p.Despawn()
	p.SendPacket(&packet.CombatDeath{
		Message:  message,
		PlayerID: p.entityID,
	})
}

func (p *PlayerController) Position() (x, y, z float64) {
	return p.player.Position()
}

func (p *PlayerController) Rotation() (yaw, pitch float32) {
	return p.player.Rotation()
}

func (p *PlayerController) OnGround() bool {
	return p.player.OnGround()
}

func (p *PlayerController) GameMode() byte {
	return p.player.GameMode()
}

func (p *PlayerController) SetGameMode(gm byte) {
	p.player.SetGameMode(gm)
	p.SendPacket(&packet.GameEvent{
		Event: 3,
		Value: float32(gm),
	})

	p.player.SetGameMode(byte(int32(gm)))
	p.Server.PlayerlistUpdate()
}

func (p *PlayerController) Push(x, y, z float64) {
	yaw, pitch := p.player.Rotation()
	p.player.SetPosition(x, y, z, yaw, pitch, p.player.OnGround())
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: idCounter.Add(1),
	})
	p.BroadcastMovement(0, x, y, z, yaw, pitch, p.player.OnGround(), true)
}

func (p *PlayerController) Teleport(x, y, z float64, yaw, pitch float32) {
	p.player.SetPosition(x, y, z, yaw, pitch, p.player.OnGround())
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: idCounter.Add(1),
	})
	p.BroadcastMovement(0, x, y, z, yaw, pitch, p.player.OnGround(), true)
}

// for now ig

func (p *PlayerController) SendCommands(g *commands.Graph) {
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

func (p *PlayerController) Keepalive() {
	id := rand.Int63() * 100
	p.SendPacket(&packet.KeepAlive{PayloadID: id})
}

func (p *PlayerController) Disconnect(reason string) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.SendPacket(pk)
	p.conn.Close(nil)
}

func distance2i(x, z int32) float64 {
	return math.Sqrt(float64(x*x) + float64(z*z))
}

func (p *PlayerController) SendChunks(dimension *world.Dimension) {
	max := int32(p.clientInfo.ViewDistance)
	px, _, pz := p.Position()
	cx, cz := int32(px)/16, int32(pz)/16

	if p.loadedChunks == nil {
		p.loadedChunks = make(map[[2]int32]struct{})
	}

	for x := -(cx + max); x <= cx+max; x++ {
		for z := -(cz + max); x <= cz+max; x++ {
			if _, ok := p.loadedChunks[[2]int32{x, z}]; ok {
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

				var e *Entity

				if f := p.Server.FindEntityByUUID(u); f != nil {
					if d, ok := f.(*Entity); ok {
						e = d
					}
				} else {
					e = p.Server.NewEntity(en)
				}

				t, ok := registry.GetEntity(e.data.Id)
				if !ok {
					continue
				}

				p.SendPacket(&packet.SpawnEntity{
					EntityID:  e.ID,
					UUID:      u,
					X:         e.data.Pos[0],
					Y:         e.data.Pos[1],
					Z:         e.data.Pos[2],
					Pitch:     degreesToAngle(e.data.Rotation[1]),
					Yaw:       degreesToAngle(e.data.Rotation[0]),
					VelocityX: int16(e.data.Motion[0]),
					VelocityY: int16(e.data.Motion[1]),
					VelocityZ: int16(e.data.Motion[2]),
					Type:      t.ProtocolID,
				})
				p.spawnedEntities = append(p.spawnedEntities, e.ID)
			}
		}
	}
}

func (p *PlayerController) SendSpawnChunks(dimension *world.Dimension) {
	max := float64(p.Server.Config.ViewDistance)
	if p.loadedChunks == nil {
		p.loadedChunks = make(map[[2]int32]struct{})
	}

	x1, _, z1 := p.player.Position()

	chunkX := math.Floor(x1 / 16)
	chunkZ := math.Floor(z1 / 16)

	for x := chunkX - max; x <= chunkX+max; x++ {
		for z := chunkZ - max; z <= chunkZ+max; z++ {
			if _, ok := p.loadedChunks[[2]int32{int32(x), int32(z)}]; ok {
				continue
			}
			c, err := dimension.Chunk(int32(x), int32(z))
			if err != nil {
				continue
			}
			p.loadedChunks[[2]int32{int32(x), int32(z)}] = struct{}{}
			p.SendPacket(c.Data())

			/*for _, en := range c.Entities {
				u, _ := world.NBTToUUID(en.UUID)

				var e *Entity

				if f := p.Server.FindEntityByUUID(u); f != nil {
					if d, ok := f.(*Entity); ok {
						e = d
					}
				} else {
					e = p.Server.NewEntity(en)
				}

				t, ok := registry.GetEntity(e.data.Id)
				if !ok {
					continue
				}

				p.SendPacket(&packet.SpawnEntity{
					EntityID:  e.ID,
					UUID:      u,
					X:         e.data.Pos[0],
					Y:         e.data.Pos[1],
					Z:         e.data.Pos[2],
					Pitch:     degreesToAngle(e.data.Rotation[1]),
					Yaw:       degreesToAngle(e.data.Rotation[0]),
					VelocityX: int16(e.data.Motion[0]),
					VelocityY: int16(e.data.Motion[1]),
					VelocityZ: int16(e.data.Motion[2]),
					Type:      t.ProtocolID,
				})
				p.spawnedEntities = append(p.spawnedEntities, e.ID)
			}*/
		}
	}
}

func (p *PlayerController) Chat(message string) {
	if !p.HasPermissions([]string{"server.chat"}) {
		return
	}
	prefix, suffix := p.GetPrefixSuffix()
	msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
		"player":        p.Name(),
		"player_prefix": prefix,
		"player_suffix": suffix,
		"message":       message,
	})

	if !p.HasPermissions([]string{"server.chat.colors"}) {
		// strip colors
		sp := strings.Split(msg, "")
		for i, c := range sp {
			if c == "&" {
				if sp[i+1] != " " {
					sp = slices.Delete(sp, i, i+2)
				}
			}
		}
		msg = strings.Join(sp, "")
	}

	p.Server.GlobalMessage(msg, p)
}

func (p *PlayerController) GetPrefixSuffix() (prefix string, suffix string) {
	group := getGroup(getPlayer(p.UUID).Group)
	return group.Prefix, group.Suffix
}

func (p *PlayerController) HandleCenterChunk(x1, z1, x2, z2 float64) {
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

func (p *PlayerController) IsSpawned(entityId int32) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, e := range p.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (p *PlayerController) SpawnPlayer(pl *PlayerController) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entityId := pl.entityID
	x, y, z := pl.player.Position()
	ya, pi := pl.player.Rotation()
	yaw, pitch := degreesToAngle(ya), degreesToAngle(pi)

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

	p.spawnedEntities = append(p.spawnedEntities, entityId)
}

func (p *PlayerController) DespawnPlayer(pl *PlayerController) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entityId := pl.entityID

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

func (p *PlayerController) InitializeInventory() {
	p.SendPacket(&SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.player.Inventory(),
	})
}

type SetContainerContent struct {
	WindowID uint8
	StateID  int32
	Slots    []world.Slot
}

func (m SetContainerContent) ID() int32 {
	return 0x12
}

func (m *SetContainerContent) Decode(r *packet.Reader) error {
	//todo reader
	return nil
}

func (m SetContainerContent) Encode(w packet.Writer) error {
	w.Uint8(m.WindowID)
	w.VarInt(m.StateID)
	if m.WindowID == 0 {
		m.Slots = sortInventory(m.Slots)
	}
	w.VarInt(int32(len(m.Slots)))
	for _, s := range m.Slots {
		i, ok := registry.GetItem(s.Id)
		if !ok {
			w.Bool(false)
			continue
		}
		w.Bool(true)
		w.VarInt(i.ProtocolID)
		w.Int8(s.Count)
		w.Int8(0)
	}
	w.Bool(false)
	return nil
}

func dataSlotToNetworkSlot(index int) int {
	switch {
	case index == 100:
		index = 8
	case index == 101:
		index = 7
	case index == 102:
		index = 6
	case index == 103:
		index = 5
	case index == -106:
		index = 45
	case index <= 8:
		index += 36
	case index >= 80 && index <= 83:
		index -= 79
	}
	return index
}

func sortInventory(slots []world.Slot) []world.Slot {
	a := make([]world.Slot, 46)
	for i, s := range slots {
		a[dataSlotToNetworkSlot(i)] = s
	}
	return a
}

func (p *PlayerController) SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch) {
	p.SendPacket(&packet.CommandSuggestionsResponse{
		TransactionId: id,
		Start:         start,
		Length:        length,
		Matches:       matches,
	})
}
