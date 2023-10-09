package server

import (
	"math"
	"math/rand"
	"slices"
	"sync"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/item"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type PlayerController struct {
	mu      sync.RWMutex
	player  *player.Player
	session *Session
	Server  *Server

	spawnedEntities []int32
	loadedChunks    map[[2]int32]struct{}

	UUID string
}

func (p *PlayerController) Name() string {
	return p.session.conn.Info.Name
}

func (p *PlayerController) Login(d *world.Dimension) error {
	if err := p.session.SendPacket(&packet.JoinGame{
		EntityID:            p.player.EntityId(),
		IsHardcore:          p.player.IsHardcore(),
		GameMode:            p.player.GameMode(),
		PreviousGameMode:    -1,
		DimensionNames:      []string{d.Type()},
		DimensionType:       d.Type(),
		DimensionName:       d.Type(),
		HashedSeed:          d.Seed(),
		MaxPlayers:          0,
		ViewDistance:        p.player.ViewDistance(),
		SimulationDistance:  p.player.SimulationDistance(),
		ReducedDebugInfo:    false,
		EnableRespawnScreen: false,
		IsDebug:             false,
		IsFlat:              false,
		DeathDimensionName:  "",
		DeathLocation:       0,
		PartialCooldown:     0,
	}); err != nil {
		return err
	}
	p.session.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite"),
	})

	x1, y1, z1 := p.player.SavedPosition()
	yaw, pitch := p.player.SavedRotation()

	chunkX, chunkZ := math.Floor(x1/16), math.Floor(z1/16)
	p.session.SendPacket(&packet.SetCenterChunk{ChunkX: int32(chunkX), ChunkZ: int32(chunkZ)})
	p.Teleport(x1, y1, z1, yaw, pitch)
	p.SendSpawnChunks()
	p.Teleport(x1, y1, z1, yaw, pitch)

	x, y, z, a := p.Server.World.Spawn()

	return p.session.SendPacket(&packet.SetDefaultSpawnPosition{
		Location: ((uint64(x) & 0x3FFFFFF) << 38) | ((uint64(z) & 0x3FFFFFF) << 12) | (uint64(y) & 0xFFF),
		Angle:    a,
	})
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.session.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) SetHealth(health float32) {
	p.player.SetHealth(health)
	food, saturation := p.player.FoodLevel(), p.player.FoodSaturationLevel()
	p.session.SendPacket(&packet.SetHealth{
		Health:         health,
		Food:           food,
		FoodSaturation: saturation,
	})
	if health == 0 {
		p.Kill("died :skull:")
	}
}

func (p *PlayerController) Kill(message string) {
	p.BroadcastHealth()
	p.BroadcastPose(7)
	if f, _ := world.GameRule(p.Server.World.Gamerules()["doImmediateRespawn"]).Bool(); !f {
		p.session.SendPacket(&packet.GameEvent{
			Event: 11,
			Value: 0,
		})
	}
	p.session.SendPacket(&packet.CombatDeath{
		Message:  message,
		PlayerID: p.player.EntityId(),
	})
}

func (p *PlayerController) ClientSettings() player.ClientInformation {
	return p.player.ClientSettings()
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
	p.session.SendPacket(&packet.GameEvent{
		Event: 3,
		Value: float32(gm),
	})
	p.session.conn.Info.GameMode = int32(gm)
	p.Server.PlayerlistUpdate()
}

func (p *PlayerController) Teleport(x, y, z float64, yaw, pitch float32) {
	p.Server.teleportCounter++
	p.player.SetPosition(x, y, z, yaw, pitch, p.player.OnGround())
	p.session.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: p.Server.teleportCounter,
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
	p.session.SendPacket(graph.Data())
}

func (p *PlayerController) Keepalive() {
	id := rand.Int63() * 100
	p.session.SendPacket(&packet.KeepAlive{PayloadID: id})
}

func (p *PlayerController) Disconnect(reason string) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.session.SendPacket(pk)
}

func distance2i(x, z int32) float64 {
	return math.Sqrt(float64(x*x) + float64(z*z))
}

func (p *PlayerController) CalculateUnusedChunks() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for c := range p.loadedChunks {
		x, _, z := p.Position()
		px, pz := int32(x)/16, int32(z)/16
		if distance2i(c[0]-px, c[1]-pz) > float64(p.ClientSettings().ViewDistance) {
			p.session.SendPacket(&packet.UnloadChunk{
				ChunkX: c[0],
				ChunkZ: c[1],
			})
			delete(p.loadedChunks, c)
		}
	}
}

func (p *PlayerController) SendChunks() {
	ow := p.Server.World.Overworld()
	max := int32(p.player.ClientSettings().ViewDistance)
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
			c, err := ow.Chunk(x, z)
			if err != nil {
				continue
			}
			p.loadedChunks[[2]int32{x, z}] = struct{}{}
			p.session.SendPacket(c.Data())
		}
	}
}

func (p *PlayerController) SendSpawnChunks() {
	ow := p.Server.World.Overworld()
	max := float64(p.Server.Config.ViewDistance)
	if p.loadedChunks == nil {
		p.loadedChunks = make(map[[2]int32]struct{})
	}

	x1, _, z1 := p.player.Position()

	chunkX := math.Abs(math.Floor(x1/16)) + max
	chunkZ := math.Abs(math.Floor(z1/16)) + max

	for x := -chunkX; x <= chunkX; x++ {
		for z := -chunkZ; z <= chunkZ; z++ {
			if _, ok := p.loadedChunks[[2]int32{int32(x), int32(z)}]; ok {
				continue
			}
			c, err := ow.Chunk(int32(x), int32(z))
			if err != nil {
				continue
			}
			p.loadedChunks[[2]int32{int32(x), int32(z)}] = struct{}{}
			p.session.SendPacket(c.Data())
		}
	}
}

func (p *PlayerController) Chat(message string) {
	prefix, suffix := p.GetPrefixSuffix()
	msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
		"player":        p.Name(),
		"player_prefix": prefix,
		"player_suffix": suffix,
		"message":       message,
	})
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
		p.session.SendPacket(&packet.SetCenterChunk{
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
	entityId := pl.player.EntityId()
	x, y, z := pl.player.Position()
	ya, pi := pl.player.Rotation()
	yaw, pitch := degreesToAngle(ya), degreesToAngle(pi)

	p.session.SendPacket(&packet.SpawnPlayer{
		EntityID:   entityId,
		PlayerUUID: pl.session.Info().UUID,
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
	})
	p.spawnedEntities = append(p.spawnedEntities, entityId)
}

func (p *PlayerController) DespawnPlayer(pl *PlayerController) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entityId := pl.player.EntityId()

	p.session.SendPacket(&packet.DestroyEntities{
		EntityIds: []int32{entityId},
	})
	index := -1
	for i, e := range p.spawnedEntities {
		if e == entityId {
			index = i
		}
	}
	if index > -1 {
		p.spawnedEntities = slices.Delete(p.spawnedEntities, index, index+1)
	}
}

func (p *PlayerController) InitializeInventory() {
	p.session.SendPacket(&SetContainerContent{
		WindowID: 0,
		StateID:  1,
		Slots:    p.player.GetSavedInventory(),
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
		i, ok := item.GetItem(s.Id)
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
	p.session.SendPacket(&packet.CommandSuggestionsResponse{
		TransactionId: id,
		Start:         start,
		Length:        length,
		Matches:       matches,
	})
}
