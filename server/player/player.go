package player

import (
	"maps"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/level"
)

var _ entity.LivingEntity = (*Player)(nil)

type Player struct {
	entityId int32

	data       level.PlayerData
	x, y, z    atomic.AtomicValue[float64]
	vx, vy, vz atomic.AtomicValue[float64]
	yaw, pitch atomic.AtomicValue[float32]
	onGround   atomic.AtomicValue[bool]

	health         atomic.AtomicValue[float32]
	food           atomic.AtomicValue[int32]
	foodExhaustion atomic.AtomicValue[float32]
	foodSaturation atomic.AtomicValue[float32]

	abilities atomic.AtomicValue[level.PlayerAbilities]

	dimension atomic.AtomicValue[string]

	gameMode atomic.AtomicValue[level.GameMode]

	selectedItemSlot atomic.AtomicValue[int32]

	recipeBook atomic.AtomicValue[level.RecipeBook]

	md_mu    sync.RWMutex
	metadata metadata.Metadata

	inventory *container.Container

	att_mu     sync.RWMutex
	attributes []entity.Attribute
}

// looks up a player in the cache or creates one if not found
func (mgr *PlayerManager) New(data level.PlayerData) *Player {
	if p, ok := mgr.lookup(data.UUID.UUID()); ok {
		return p
	}

	pl := &Player{
		entityId: entity.NewEntityId(),
		metadata: metadata.Metadata{
			// Entity
			metadata.BaseIndex:                      metadata.Byte(0),
			metadata.AirTicksIndex:                  metadata.VarInt(300),
			metadata.CustomNameIndex:                metadata.OptionalTextComponent(nil),
			metadata.IsCustomNameVisibleIndex:       metadata.Boolean(false),
			metadata.IsSilentIndex:                  metadata.Boolean(false),
			metadata.HasNoGravityIndex:              metadata.Boolean(false),
			metadata.PoseIndex:                      metadata.Standing,
			metadata.TicksFrozenInPowderedSnowIndex: metadata.VarInt(0),
			// Living Entity extends Entity
			metadata.LivingEntityHandstatesIndex: metadata.Byte(0),
			metadata.LivingEntityHealthIndex:     metadata.Float(data.Health),
			//metadata.LivingEntityPotionEffectColorIndex:   metadata.VarInt(0),
			metadata.LivingEntityPotionEffectAmbientIndex: metadata.Boolean(false),
			metadata.LivingEntityArrowCountIndex:          metadata.VarInt(0),
			metadata.LivingEntityBeeStingersCountIndex:    metadata.VarInt(0),
			metadata.LivingEntitySleepingBedPositionIndex: metadata.OptionalPosition(nil),
			// Player extends Living Entity
			metadata.PlayerAdditionalHeartsIndex:   metadata.Float(0),
			metadata.PlayerScoreIndex:              metadata.VarInt(0),
			metadata.PlayerDisplayedSkinPartsIndex: metadata.Byte(0),
			metadata.PlayerMainHandIndex:           metadata.Byte(1),
		},

		x: atomic.Value(data.Pos[0]),
		y: atomic.Value(data.Pos[1]),
		z: atomic.Value(data.Pos[2]),

		vx: atomic.Value(data.Motion[0]),
		vy: atomic.Value(data.Motion[1]),
		vz: atomic.Value(data.Motion[2]),

		yaw:   atomic.Value(data.Rotation[0]),
		pitch: atomic.Value(data.Rotation[1]),

		onGround: atomic.Value(data.OnGround),

		dimension: atomic.Value(data.Dimension),

		gameMode: atomic.Value(data.PlayerGameType),

		recipeBook: atomic.Value(data.RecipeBook),

		selectedItemSlot: atomic.Value(data.SelectedItemSlot),

		health:         atomic.Value(data.Health),
		food:           atomic.Value(data.FoodLevel),
		foodExhaustion: atomic.Value(data.FoodExhaustionLevel),
		foodSaturation: atomic.Value(data.FoodSaturationLevel),

		abilities: atomic.Value(data.Abilities),

		inventory: &data.Inventory,

		attributes: data.Attributes,

		data: data,
	}
	mgr.add(pl)

	return pl
}

func (p *Player) Type() int32 {
	return registry.EntityType.Get("minecraft:player")
}

func (p *Player) UUID() uuid.UUID {
	return p.data.UUID.UUID()
}

func (p *Player) Position() (x, y, z float64) {
	return p.x.Get(), p.y.Get(), p.z.Get()
}

func (p *Player) Rotation() (yaw, pitch float32) {
	return p.yaw.Get(), p.pitch.Get()
}

func (p *Player) OnGround() bool {
	return p.onGround.Get()
}

func (p *Player) SetPosition(x, y, z float64) {
	p.x.Set(x)
	p.y.Set(y)
	p.z.Set(z)
}

func (p *Player) SetRotation(yaw, pitch float32) {
	p.yaw.Set(yaw)
	p.pitch.Set(pitch)
}

func (p *Player) SetOnGround(val bool) {
	p.onGround.Set(val)
}

func (p *Player) Motion() (x, y, z float64) {
	return p.vx.Get(), p.vy.Get(), p.vz.Get()
}

func (p *Player) SetMotion(x, y, z float64) {
	p.vx.Set(x)
	p.vy.Set(y)
	p.vz.Set(z)
}

func (p *Player) EntityId() int32 {
	return p.entityId
}

// returns a clone of the metadata of this player
func (p *Player) Metadata() metadata.Metadata {
	p.md_mu.RLock()
	defer p.md_mu.RUnlock()
	return maps.Clone(p.metadata)
}

func (p *Player) SetMetadata(md metadata.Metadata) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	p.metadata = md
}

func (p *Player) MetadataIndex(i byte) any {
	p.md_mu.RLock()
	defer p.md_mu.RUnlock()
	return p.metadata[i]
}

func (p *Player) SetMetadataIndex(i byte, v any) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	p.metadata[i] = v
}

func (p *Player) SetMetadataIndexes(md metadata.Metadata) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	for index, value := range md {
		p.metadata[index] = value
	}
}

func (p *Player) Dimension() string {
	return p.dimension.Get()
}

func (p *Player) SetDimension(dim string) {
	p.dimension.Set(dim)
}

func (p *Player) Health() float32 {
	return p.health.Get()
}

func (p *Player) SetHealth(h float32) {
	p.health.Set(h)
}

func (p *Player) Food() int32 {
	return p.food.Get()
}

func (p *Player) SetFood(f int32) {
	p.food.Set(f)
}

func (p *Player) FoodSaturation() float32 {
	return p.foodSaturation.Get()
}

func (p *Player) SetFoodSaturation(fs float32) {
	p.foodSaturation.Set(fs)
}

func (p *Player) FoodExhaustion() float32 {
	return p.foodExhaustion.Get()
}

func (p *Player) SetFoodExhaustion(fh float32) {
	p.foodExhaustion.Set(fh)
}

func (p *Player) Abilities() level.PlayerAbilities {
	return p.abilities.Get()
}

func (p *Player) SetAbilities(abs level.PlayerAbilities) {
	p.abilities.Set(abs)
}

func (p *Player) GameMode() level.GameMode {
	return p.gameMode.Get()
}

func (p *Player) SetGameMode(mode level.GameMode) {
	p.gameMode.Set(mode)
}

func (p *Player) Attribute(id string) *entity.Attribute {
	p.att_mu.RLock()
	defer p.att_mu.RUnlock()
	i := slices.IndexFunc(p.attributes, func(att entity.Attribute) bool { return att.Id == id })
	if i == -1 {
		return nil
	}
	return &p.attributes[i]
}

// returns a clone of the attributes of this player
func (p *Player) Attributes() []entity.Attribute {
	p.att_mu.RLock()
	defer p.att_mu.RUnlock()
	return slices.Clone(p.attributes)
}

func (p *Player) SetAttribute(id string, base float64) {
	p.att_mu.Lock()
	defer p.att_mu.Unlock()
	i := slices.IndexFunc(p.attributes, func(att entity.Attribute) bool { return att.Id == id })
	if i == -1 {
		return
	}
	p.attributes[i].Base = base
}

func (p *Player) RecipeBook() level.RecipeBook {
	return p.recipeBook.Get()
}

func (p *Player) SetRecipeBook(book level.RecipeBook) {
	p.recipeBook.Set(book)
}

func (p *Player) Inventory() *container.Container {
	return p.inventory
}

// if negative, returns 0 and if over 8, returns 8
func (p *Player) SelectedItemSlot() int32 {
	slot := p.selectedItemSlot.Get()
	if slot < 0 {
		slot = 0
	}
	if slot > 8 {
		slot = 8
	}
	return slot
}

// if negative, set to 0 and if over 8, set to 8
func (p *Player) SetSelectedItemSlot(slot int32) {
	if slot < 0 {
		slot = 0
	}
	if slot > 8 {
		slot = 8
	}
	p.selectedItemSlot.Set(slot)
}

func (p *Player) sync() {
	x, y, z := p.Position()
	yaw, pitch := p.Rotation()

	p.data.Abilities = p.abilities.Get()
	p.data.Pos = [3]float64{x, y, z}
	p.data.Rotation = [2]float32{yaw, pitch}
	p.data.OnGround = p.onGround.Get()
	p.data.Dimension = p.dimension.Get()
	p.data.Inventory = *p.inventory
	p.data.RecipeBook = p.recipeBook.Get()

	p.att_mu.RLock()
	p.data.Attributes = p.attributes
	p.att_mu.RUnlock()

	p.data.Health = p.health.Get()
	p.data.FoodLevel = p.food.Get()
	p.data.FoodExhaustionLevel = p.foodExhaustion.Get()
	p.data.FoodSaturationLevel = p.foodSaturation.Get()
	p.data.PlayerGameType = p.gameMode.Get()
	p.data.SelectedItemSlot = p.selectedItemSlot.Get()

	//TODO motion(velocity), xp, etc
}
