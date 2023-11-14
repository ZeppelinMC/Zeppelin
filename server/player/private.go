package player

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/item"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/google/uuid"
)

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

func (p *Player) ClientSettings() clientInfo {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.clientInfo
}

func (p *Player) SetClientSettings(pk *packet.ClientSettings) {
	p.mu.Lock()
	p.clientInfo.Locale = pk.Locale
	//don't set view distance but server controls it
	p.clientInfo.ChatMode = pk.ChatMode
	p.clientInfo.ChatColors = pk.ChatColors
	p.clientInfo.DisplayedSkinParts = pk.DisplayedSkinParts
	p.clientInfo.MainHand = pk.MainHand
	p.clientInfo.DisableTextFiltering = pk.DisableTextFiltering
	p.clientInfo.AllowServerListings = pk.AllowServerListings
	p.mu.Unlock()

	p.BroadcastMetadataInArea(&packet.SetEntityMetadata{
		DisplayedSkinParts: &pk.DisplayedSkinParts,
	})
}

func (p *Player) Properties() []types.Property {
	return p.conn.Properties()
}

func (p *Player) Dimension() *world.Dimension {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.dimension
}

func (p *Player) SetDimension(d *world.Dimension) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dimension = d
}

func (p *Player) IsDead() bool {
	return p.dead.Load()
}

func (p *Player) SetDead(a bool) {
	p.dead.Store(a)
}

func (p *Player) Health() float32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) SetHealth(health float32) {
	if health < 0 {
		health = 0
	}
	p.mu.Lock()
	p.health = health
	p.mu.Unlock()
	food, saturation := p.FoodLevel(), p.FoodSaturationLevel()
	p.SendPacket(&packet.SetHealth{
		Health:         health,
		Food:           food,
		FoodSaturation: saturation,
	})
	/*p.broadcastMetadataGlobal(&packet.SetEntityMetadata{
		EntityID: p.EntityID(),
		Health:   &health,
	})*/
	if health <= 0 {
		p.Kill("died :skull:")
	}
}

func (p *Player) FoodLevel() int32 {
	return p.food.Load()
}

func (p *Player) SetFoodLevel(level int32) {
	p.food.Store(level)
}

func (p *Player) FoodSaturationLevel() float32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.foodSaturation
}

func (p *Player) SetFoodSaturationLevel(level float32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.foodSaturation = level
}

func (p *Player) SavedAbilities() world.Abilities {
	return p.data.Abilities
}

func (p *Player) SetFlying(val bool) {
	p.flying.Store(val)
}

func (p *Player) IsHardcore() bool {
	return p.isHardCore.Load()
}

func (p *Player) SetGameMode(gm byte) {
	p.mu.Lock()
	p.gameMode = gm
	p.mu.Unlock()
	p.SendPacket(&packet.GameEvent{
		Event: enum.GameEventChangeGamemode,
		Value: float32(gm),
	})
	p.BroadcastGamemode()
}

func (p *Player) GameMode() byte {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.gameMode
}

func (p *Player) Position() (x, y, z float64) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.x, p.y, p.z
}

func (p *Player) Rotation() (yaw, pitch float32) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.yaw, p.pitch
}

func (p *Player) OnGround() bool {
	return p.onGround.Load()
}

func (p *Player) SetPosition(x, y, z float64, yaw, pitch float32, ong bool) {
	p.onGround.Store(ong)
	p.mu.Lock()
	defer p.mu.Unlock()
	p.x, p.y, p.z, p.yaw, p.pitch = x, y, z, yaw, pitch
}

func (p *Player) SetHighestY(y float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.highestY = y
}

func (p *Player) HighestY() float64 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.highestY
}

func (p *Player) Operator() bool {
	return p.operator.Load()
}

func (p *Player) SetOperator(op bool) {
	p.operator.Store(op)
	v := enum.EntityStatusPlayerOpPermissionLevel0
	if op {
		v = enum.EntityStatusPlayerOpPermissionLevel4
	}
	p.SendPacket(&packet.EntityEvent{
		EntityID: p.entityID,
		Status:   v,
	})
}

func (p *Player) SetSelectedSlot(h int32) {
	p.selectedSlot.Store(h)
	p.Inventory.SetSelectedSlot(h)
}

func (p *Player) SelectedSlot() int32 {
	return p.selectedSlot.Load()
}

func (p *Player) PreviousSelectedSlot() item.Item {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.previousSelectedSlot
}

func (p *Player) SetPreviousSelectedSlot(s item.Item) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.previousSelectedSlot = s
}

func (p *Player) SetSessionID(id [16]byte, pk, ks []byte, expires int64) {
	p.mu.Lock()
	p.sessionID = id
	p.publicKey = bytes.Clone(pk)
	p.keySignature = bytes.Clone(ks)
	p.expires = expires
	p.mu.Unlock()

	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x02,
			Players: []types.PlayerInfo{
				{
					UUID:          p.conn.UUID(),
					ChatSessionID: id,
					PublicKey:     bytes.Clone(pk),
					KeySignature:  bytes.Clone(ks),
					ExpiresAt:     expires,
				},
			},
		})
		return true
	})
}

func (p *Player) SessionID() (id [16]byte, pk, ks []byte, expires int64) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.sessionID, p.publicKey, p.keySignature, p.expires
}

// SetSkin allows you to set the player's skin.
func (p *Player) SetSkin(url string) {
	var textures types.TexturesProperty
	b, _ := base64.StdEncoding.DecodeString(p.conn.Properties()[0].Value)
	json.Unmarshal(b, &textures)
	textures.Textures.Skin.URL = url

	t, _ := json.Marshal(textures)

	d := base64.StdEncoding.EncodeToString(t)

	p.conn.Properties()[0].Signature = ""
	p.conn.Properties()[0].Value = d
}

// SetCape allows you to set the player's cape.
func (p *Player) SetCape(url string) {
	var textures types.TexturesProperty
	b, _ := base64.StdEncoding.DecodeString(p.conn.Properties()[0].Value)
	json.Unmarshal(b, &textures)
	textures.Textures.Cape.URL = url

	t, _ := json.Marshal(textures)

	d := base64.StdEncoding.EncodeToString(t)

	p.conn.Properties()[0].Signature = ""
	p.conn.Properties()[0].Value = d
}

func (p *Player) SetDisplayName(name *chat.Message) {
	p.playerController.Range(func(_ uuid.UUID, pl *Player) bool {
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x20,
			Players: []types.PlayerInfo{
				{
					UUID:        p.UUID(),
					DisplayName: name,
				},
			},
		})
		return true
	})
	p.BroadcastMetadataInArea(&packet.SetEntityMetadata{
		EntityID:            p.entityID,
		CustomName:          name,
		IsCustomNameVisible: point(name != nil),
	})
}
