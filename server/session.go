package server

import (
	"math"
	"math/rand"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/dynamitemc/dynamite/server/network/handlers"

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

type Session struct {
	mu        sync.RWMutex
	player    *player.Player
	conn      *minecraft.Conn
	Server    *Server
	sessionID [16]byte
	publicKey []byte

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

	//playReady means the player is ready to receive packets regarding other players
	playReady bool
}

func (p *Session) HandlePackets() error {
	p.playReady = true
	ticker := time.NewTicker(25 * time.Second)
	for {
		select {
		case <-ticker.C:
			p.Keepalive()
		default:
		}

		packt, err := p.conn.ReadPacket()
		if err != nil {
			return err
		}

		switch pk := packt.(type) {
		case *packet.PlayerCommandServer:
			handlers.PlayerCommand(p, pk.ActionID)
		case *packet.ChatMessageServer:
			handlers.ChatMessagePacket(p, pk)
		case *packet.ChatCommandServer:
			handlers.ChatCommandPacket(p, p.Server.commandGraph, pk.Command)
		case *packet.ClientSettings:
			handlers.ClientSettings(p, pk)
		case *packet.PlayerPosition, *packet.PlayerPositionRotation, *packet.PlayerRotation:
			handlers.PlayerMovement(p, p.player, pk)
		case *packet.PlayerActionServer:
			handlers.PlayerAction(p, pk)
		case *packet.InteractServer:
			handlers.Interact(p, pk)
		case *packet.SwingArmServer:
			handlers.SwingArm(p, pk.Hand)
		case *packet.CommandSuggestionsRequest:
			handlers.CommandSuggestionsRequest(pk.TransactionId, pk.Text, p.Server.commandGraph, p)
		case *packet.ClientCommandServer:
			handlers.ClientCommand(p, p.player, pk.ActionID)
		case *packet.PlayerAbilitiesServer:
			handlers.PlayerAbilities(p.player, pk.Flags)
		case *packet.PlayerSessionServer:
			handlers.PlayerSession(p, pk.SessionID, pk.PublicKey.PublicKey)
		}
	}
}

func (p *Session) Name() string {
	return p.conn.Name()
}

func (p *Session) SendPacket(pk packet.Packet) error {
	return p.conn.SendPacket(pk)
}

func (p *Session) SetClientSettings(pk *packet.ClientSettings) {
	p.clientInfo.Locale = pk.Locale
	//don't set view distance but server controls it
	p.clientInfo.ChatMode = pk.ChatMode
	p.clientInfo.ChatColors = pk.ChatColors
	p.clientInfo.DisplayedSkinParts = pk.DisplayedSkinParts
	p.clientInfo.MainHand = pk.MainHand
	p.clientInfo.DisableTextFiltering = pk.DisableTextFiltering
	p.clientInfo.AllowServerListings = pk.AllowServerListings
}

func (p *Session) Respawn(dim string) {
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

	clear(p.spawnedEntities)

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

func (p *Session) Login(dim string) error {
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

func (p *Session) SystemChatMessage(s string) error {
	return p.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *Session) SetHealth(health float32) {
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

func (p *Session) Kill(message string) {
	p.player.SetDead(true)
	p.BroadcastHealth()
	if f, _ := world.GameRule(p.Server.World.Gamerules()["doImmediateRespawn"]).Bool(); !f {
		p.SendPacket(&packet.GameEvent{
			Event: 11,
			Value: 0,
		})
	}

	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()

	for _, pl := range p.Server.Players {
		if !p.IsSpawned(p.entityID) {
			continue
		}
		pl.SendPacket(&packet.DamageEvent{
			EntityID:     p.entityID,
			SourceTypeID: 0,
		})
	}

	p.SendPacket(&packet.DamageEvent{
		EntityID:     p.entityID,
		SourceTypeID: 0,
	})
	p.Despawn()
	p.SendPacket(&packet.CombatDeath{
		Message:  message,
		PlayerID: p.entityID,
	})
}

func (p *Session) SetSessionID(id [16]byte, pk []byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.sessionID = id
	p.publicKey = pk
}

func (p *Session) Position() (x, y, z float64) {
	return p.player.Position()
}

func (p *Session) Rotation() (yaw, pitch float32) {
	return p.player.Rotation()
}

func (p *Session) OnGround() bool {
	return p.player.OnGround()
}

func (p *Session) GameMode() byte {
	return p.player.GameMode()
}

func (p *Session) SetGameMode(gm byte) {
	p.player.SetGameMode(gm)
	p.SendPacket(&packet.GameEvent{
		Event: 3,
		Value: float32(gm),
	})

	p.player.SetGameMode(byte(int32(gm)))
	p.Server.PlayerlistUpdate()
}

func (p *Session) Push(x, y, z float64) {
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

func (p *Session) Teleport(x, y, z float64, yaw, pitch float32) {
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

func (p *Session) SendCommands(g *commands.Graph) {
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

func (p *Session) Keepalive() {
	id := rand.Int63() * 100
	p.SendPacket(&packet.KeepAlive{PayloadID: id})
}

func (p *Session) Disconnect(reason string) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.SendPacket(pk)
	p.conn.Close(nil)
}

func distance2i(x, z int32) float64 {
	return math.Sqrt(float64(x*x) + float64(z*z))
}

func (p *Session) SendChunks(dimension *world.Dimension) {
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
				p.mu.Lock()
				p.spawnedEntities = append(p.spawnedEntities, e.ID)
				p.mu.Unlock()
			}
		}
	}
}

func (p *Session) SendSpawnChunks(dimension *world.Dimension) {
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
				p.mu.Lock()
				p.spawnedEntities = append(p.spawnedEntities, e.ID)
				p.mu.Unlock()
			}
		}
	}
}

func (p *Session) Chat(pk *packet.ChatMessageServer) {
	if !p.Server.Config.Chat.Enable {
		return
	}
	if !p.HasPermissions([]string{"server.chat"}) {
		return
	}

	if !p.Server.Config.Chat.Secure {
		prefix, suffix := p.GetPrefixSuffix()
		msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
			"player":        p.Name(),
			"player_prefix": prefix,
			"player_suffix": suffix,
			"message":       pk.Message,
		})

		if !p.Server.Config.Chat.Enable || !p.HasPermissions([]string{"server.chat.colors"}) {
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
	} else {
		p.Server.GlobalBroadcast(&packet.PlayerChatMessage{
			Sender: p.conn.UUID(),
			//MessageSignature: pk.Signature,
			Message:     pk.Message,
			Timestamp:   pk.Timestamp,
			Salt:        pk.Salt,
			NetworkName: p.Name(),
		})
	}
}

func (p *Session) GetPrefixSuffix() (prefix string, suffix string) {
	group := getGroup(getPlayer(p.UUID).Group)
	return group.Prefix, group.Suffix
}

func (p *Session) HandleCenterChunk(x1, z1, x2, z2 float64) {
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

func (p *Session) IsSpawned(entityId int32) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, e := range p.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (p *Session) SpawnPlayer(pl *Session) {
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

func (p *Session) DespawnPlayer(pl *Session) {
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

func (p *Session) InitializeInventory() {
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

func (p *Session) SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch) {
	p.SendPacket(&packet.CommandSuggestionsResponse{
		TransactionId: id,
		Start:         start,
		Length:        length,
		Matches:       matches,
	})
}
