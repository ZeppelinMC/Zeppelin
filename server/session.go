package server

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sync"
	"time"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/protocol/types"
	"github.com/google/uuid"

	"github.com/dynamitemc/dynamite/server/inventory"
	"github.com/dynamitemc/dynamite/server/item"
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
	mu     sync.RWMutex
	Player *player.Player
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

	sessionID    [16]byte
	publicKey    []byte
	keySignature []byte
	expires      uint64

	acknowledgedMessages []packet.PreviousMessage
	index                int32
}

func (p *Session) HandlePackets() error {
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
			handlers.ChatCommandPacket(p, p.Server.commandGraph, pk.Command, pk.Timestamp, pk.Salt, pk.ArgumentSignatures)
		case *packet.ClientSettings:
			handlers.ClientSettings(p, pk)
		case *packet.PlayerPosition, *packet.PlayerPositionRotation, *packet.PlayerRotation:
			handlers.PlayerMovement(p, p.Player, pk)
		case *packet.PlayerActionServer:
			handlers.PlayerAction(p, p.Player, pk)
		case *packet.InteractServer:
			handlers.Interact(p, pk)
		case *packet.SwingArmServer:
			handlers.SwingArm(p, pk.Hand)
		case *packet.CommandSuggestionsRequest:
			handlers.CommandSuggestionsRequest(pk.TransactionId, pk.Text, p.Server.commandGraph, p)
		case *packet.ClientCommandServer:
			handlers.ClientCommand(p, p.Player, pk.ActionID)
		case *packet.PlayerAbilitiesServer:
			handlers.PlayerAbilities(p.Player, pk.Flags)
		case *packet.PlayerSessionServer:
			handlers.PlayerSession(p, pk.SessionID, pk.PublicKey, pk.KeySignature, pk.ExpiresAt)
		case *packet.SetHeldItemServer:
			handlers.SetHeldItem(p.Player, pk.Slot)
		case *packet.SetCreativeModeSlot:
			handlers.SetCreativeModeSlot(p, p.Player, int8(inventory.NetworkSlotToDataSlot(pk.Slot)), pk.ClickedItem)
		case *packet.TeleportToEntityServer:
			handlers.TeleportToEntity(p, p.Player, pk.Player)
		case *packet.ClickContainer:
			handlers.ClickContainer(p, p.Player, pk)
		case *packet.MessageAcknowledgment:
			fmt.Println(pk.MessageCount)
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
	p.mu.Lock()
	defer p.mu.Unlock()
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
	p.Player.SetDead(false)
	p.SendPacket(&packet.Respawn{
		GameMode:         p.Player.GameMode(),
		PreviousGameMode: -1,
		DimensionType:    d.Type(),
		DimensionName:    d.Type(),
		HashedSeed:       d.Seed(),
	})
	p.Player.SetDimension(d.Type())

	var x1, y1, z1 int32
	var a float32

	switch d.Type() {
	case "minecraft:overworld":
		x1, y1, z1, a = p.Server.World.Spawn()
	case "minecraft:the_nether":
		x, y, z := p.Player.Position()
		x1, y1, z1 = int32(x)/8, int32(y)/8, int32(z)/8
	}

	clear(p.spawnedEntities)

	yaw, pitch := p.Player.Rotation()

	if b, _ := world.GameRule(p.Server.World.Gamerules()["keepInventory"]).Bool(); !b {
		p.Player.Inventory().Clear()
	}

	p.SendPacket(&packet.SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.Player.Inventory().Packet(),
	})

	chunkX, chunkZ := math.Floor(float64(x1)/16), math.Floor(float64(z1)/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)
	p.SendChunks(d)

	p.Teleport(float64(x1), float64(y1), float64(z1), yaw, pitch)

	p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: ((uint64(x1) & 0x3FFFFFF) << 38) | ((uint64(z1) & 0x3FFFFFF) << 12) | (uint64(y1) & 0xFFF),
		Angle:    a,
	})
}

func (p *Session) Login(dim string) {
	d := p.Server.GetDimension(dim)
	p.SendPacket(&packet.JoinGame{
		EntityID:           p.entityID,
		IsHardcore:         p.Player.IsHardcore(),
		GameMode:           p.Player.GameMode(),
		PreviousGameMode:   -1,
		DimensionNames:     []string{"minecraft:overworld", "minecraft:the_nether", "minecraft:the_end"},
		DimensionType:      d.Type(),
		DimensionName:      d.Type(),
		HashedSeed:         d.Seed(),
		MaxPlayers:         0,
		ViewDistance:       int32(p.clientInfo.ViewDistance),
		SimulationDistance: int32(p.clientInfo.ViewDistance), //todo fix this
	})
	p.Player.SetDimension(d.Type())
	p.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite"),
	})

	x1, y1, z1 := p.Player.Position()
	yaw, pitch := p.Player.Rotation()

	chunkX, chunkZ := math.Floor(x1/16), math.Floor(z1/16)
	p.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(x1, y1, z1, yaw, pitch)
	p.SendChunks(d)

	abs := p.Player.SavedAbilities()
	abps := &packet.PlayerAbilities{FlyingSpeed: abs.FlySpeed, FieldOfViewModifier: 0.1}
	if abs.Flying != 0 {
		abps.Flags |= 0x06
	}
	if abps.Flags != 0 {
		p.SendPacket(abps)
	}
	p.SendPacket(tags)

	if p.Player.Operator() {
		p.SendPacket(&packet.EntityEvent{
			EntityID: p.entityID,
			Status:   28,
		})
	}

	p.Teleport(x1, y1, z1, yaw, pitch)

	x, y, z, a := p.Server.World.Spawn()

	p.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: ((uint64(x) & 0x3FFFFFF) << 38) | ((uint64(z) & 0x3FFFFFF) << 12) | (uint64(y) & 0xFFF),
		Angle:    a,
	})

	if p.Server.Config.ResourcePack.Enable {
		p.SendPacket(&packet.ResourcePack{
			URL:    p.Server.Config.ResourcePack.URL,
			Hash:   p.Server.Config.ResourcePack.Hash,
			Forced: p.Server.Config.ResourcePack.Force,
			//Prompt: p.Server.Config.Messages.ResourcePackPrompt,
		})
	}
}

func (p *Session) SystemChatMessage(message chat.Message) error {
	return p.SendPacket(&packet.SystemChatMessage{Message: message})
}

func (p *Session) SetHealth(health float32) {
	p.Player.SetHealth(health)
	food, saturation := p.Player.FoodLevel(), p.Player.FoodSaturationLevel()
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
	p.Player.SetDead(true)
	p.BroadcastHealth()
	if f, _ := world.GameRule(p.Server.World.Gamerules()["doImmediateRespawn"]).Bool(); !f {
		p.SendPacket(&packet.GameEvent{
			Event: 11,
			Value: 0,
		})
	}

	p.Server.mu.RLock()
	for _, pl := range p.Server.players {
		if !p.IsSpawned(p.entityID) {
			continue
		}
		pl.SendPacket(&packet.DamageEvent{
			EntityID:     p.entityID,
			SourceTypeID: 0,
		})
		pl.SendPacket(&packet.EntityEvent{
			EntityID: p.entityID,
			Status:   3,
		})
	}
	p.Server.mu.RUnlock()

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

func (p *Session) SetSessionID(id [16]byte, pk, ks []byte, expires int64) {
	fmt.Println("resetting session, expires in", (expires-time.Now().UnixMilli())/1000/1000)
	p.mu.Lock()
	p.sessionID = id
	p.publicKey = pk
	p.keySignature = ks
	p.expires = uint64(expires)
	p.mu.Unlock()

	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.players {
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x02,
			Players: []types.PlayerInfo{
				{
					UUID:          p.conn.UUID(),
					ChatSessionID: id,
					PublicKey:     pk,
					KeySignature:  ks,
					ExpiresAt:     expires,
				},
			},
		})
	}
}

func (p *Session) SetGameMode(gm byte) {
	p.Player.SetGameMode(gm)
	p.SendPacket(&packet.GameEvent{
		Event: 3,
		Value: float32(gm),
	})

	p.Player.SetGameMode(byte(int32(gm)))
	p.BroadcastGamemode()
}

func (p *Session) Push(x, y, z float64) {
	yaw, pitch := p.Player.Rotation()
	p.Player.SetPosition(x, y, z, yaw, pitch, p.Player.OnGround())
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: idCounter.Add(1),
	})
	p.BroadcastMovement(0, x, y, z, yaw, pitch, p.Player.OnGround(), true)
}

func (p *Session) Teleport(x, y, z float64, yaw, pitch float32) {
	p.Player.SetPosition(x, y, z, yaw, pitch, p.Player.OnGround())
	p.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: idCounter.Add(1),
	})
	p.BroadcastMovement(0, x, y, z, yaw, pitch, p.Player.OnGround(), true)
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
	p.SendPacket(&packet.KeepAliveClient{PayloadID: id})
}

func (p *Session) Disconnect(reason chat.Message) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.SendPacket(pk)
	p.conn.Close(nil)
}

func (p *Session) IsChunkLoaded(x, z int32) bool {
	_, ok := p.loadedChunks[[2]int32{x, z}]
	return ok
}

func (p *Session) SendChunks(dimension *world.Dimension) {
	p.mu.Lock()
	defer p.mu.Unlock()
	max := float64(p.Server.Config.ViewDistance)
	if p.loadedChunks == nil {
		p.loadedChunks = make(map[[2]int32]struct{})
	}

	x1, _, z1 := p.Player.Position()

	chunkX := math.Floor(x1 / 16)
	chunkZ := math.Floor(z1 / 16)

	for x := chunkX - max; x <= chunkX+max; x++ {
		for z := chunkZ - max; z <= chunkZ+max; z++ {
			if p.IsChunkLoaded(int32(x), int32(z)) {
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
				p.spawnedEntities = append(p.spawnedEntities, e.ID)
			}
		}
	}
}

func (p *Session) UnloadChunks() {
	p.mu.Lock()
	defer p.mu.Unlock()
	x, z := p.ChunkPosition()
	for c := range p.loadedChunks {
		if d := math.Sqrt(float64((x-c[0])*(x-c[0])) + float64((z-c[1])*(z-c[1]))); d > float64(p.clientInfo.ViewDistance) {
			fmt.Println("dist", d, "vd", p.clientInfo.ViewDistance)
			fmt.Println("unloading chunk", x, z)
			p.SendPacket(&packet.UnloadChunk{
				ChunkX: x,
				ChunkZ: z,
			})
			delete(p.loadedChunks, c)
		}
	}
}

func (p *Session) ChunkPosition() (int32, int32) {
	x, _, z := p.Player.Position()
	return int32(x) / 16, int32(z) / 16
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
	p.mu.Lock()
	defer p.mu.Unlock()
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
	x, y, z := pl.Player.Position()
	ya, pi := pl.Player.Rotation()
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

	pl.SendEquipment(p)
	p.SendEquipment(pl)
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

func (p *Session) intitializeData() {
	p.SendPacket(&packet.SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.Player.Inventory().Packet(),
	})
	p.SendPacket(&packet.SetHeldItem{Slot: int8(p.Player.SelectedSlot())})
	p.SendPacket(&packet.SetHealth{
		Health:         p.Player.Health(),
		FoodSaturation: p.Player.FoodSaturationLevel(),
		Food:           p.Player.FoodLevel(),
	})
}

func (p *Session) ClearItem(slot int8) {
	p.SendPacket(&packet.SetContainerSlot{
		WindowID: 0,
		StateID:  1,
		Slot:     int16(inventory.DataSlotToNetworkSlot(slot)),
	})
	p.Player.Inventory().DeleteSlot(slot)
}

func (p *Session) SetSlot(slot int8, data packet.Slot) {
	p.SendPacket(&packet.SetContainerSlot{
		WindowID: 0,
		StateID:  1,
		Slot:     int16(inventory.DataSlotToNetworkSlot(slot)),
		Data:     data,
	})
	p.Player.Inventory().SetSlot(slot, item.Item{
		Count: data.Count,
		Slot:  int8(slot),
		Id:    registry.FindItem(data.Id),
	})
}

func (p *Session) DropSlot() {
	item := p.Player.PreviousSelectedSlot()
	x, y, z := p.Player.Position()

	id := idCounter.Add(1)
	uuid := uuid.New()

	p.Server.mu.Lock()
	defer p.Server.mu.Unlock()
	for _, pl := range p.Server.players {
		if !pl.InView(p) {
			return
		}
		pl.SendPacket(&packet.SpawnEntity{
			EntityID: id,
			UUID:     uuid,
			Type:     54,
			X:        x,
			Y:        y,
			Z:        z,
		})
		pl.SendPacket(&PacketSetPlayerMetadata{
			EntityID: id,
			Slot:     &item,
		})
	}
}
func (p *Session) SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch) {
	p.SendPacket(&packet.CommandSuggestionsResponse{
		TransactionId: id,
		Start:         start,
		Length:        length,
		Matches:       matches,
	})
}

func (p *Session) SetDisplayName(name string) {
	p.Server.mu.RLock()
	defer p.Server.mu.RUnlock()
	for _, pl := range p.Server.players {
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x20,
			Players: []types.PlayerInfo{
				{
					UUID:           p.conn.UUID(),
					HasDisplayName: name != "",
					DisplayName:    name,
				},
			},
		})
	}
}

func (p *Session) TeleportToEntity(uuid [16]byte) {
	e := p.Server.FindEntityByUUID(uuid)
	if e == nil {
		return
	}
	if pl, ok := e.(*Session); ok {
		x, y, z := pl.Player.Position()
		yaw, pitch := pl.Player.Rotation()
		p.Teleport(x, y, z, yaw, pitch)
	} else {
		e := e.(*Entity)
		x, y, z := e.data.Pos[0], e.data.Pos[1], e.data.Pos[2]
		yaw, pitch := e.data.Rotation[0], e.data.Rotation[1]
		p.Teleport(x, y, z, yaw, pitch)
	}
}
