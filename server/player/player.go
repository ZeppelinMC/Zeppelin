package player

import (
	"maps"
	"slices"
	"sync"
	a "sync/atomic"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/util/atomic"
)

var _ entity.LivingEntity = (*Player)(nil)

type Player struct {
	entityId int32

	data level.Player
	x, y, z,
	vx, vy, vz a.Uint64
	yaw, pitch a.Uint32
	onGround   a.Bool

	food a.Int32
	health,
	foodExhaustion,
	foodSaturation a.Uint32

	abilities atomic.AtomicValue[level.PlayerAbilities]

	dimension atomic.AtomicValue[string]

	gameMode         a.Int32
	selectedItemSlot a.Int32

	recipeBook atomic.AtomicValue[level.RecipeBook]

	md_mu    sync.RWMutex
	metadata metadata.Metadata

	inventory *container.Container

	att_mu     sync.RWMutex
	attributes []entity.Attribute
}

// looks up a player in the cache or creates one if not found
func (mgr *PlayerManager) New(data level.Player) *Player {
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

		x: atomicFloat64(data.Pos[0]),
		y: atomicFloat64(data.Pos[1]),
		z: atomicFloat64(data.Pos[2]),

		vx: atomicFloat64(data.Motion[0]),
		vy: atomicFloat64(data.Motion[1]),
		vz: atomicFloat64(data.Motion[2]),

		yaw:   atomicFloat32(data.Rotation[0]),
		pitch: atomicFloat32(data.Rotation[1]),

		onGround: atomicBool(data.OnGround),

		dimension: atomic.Value(data.Dimension),

		gameMode: atomicInt32(int32(data.PlayerGameType)),

		recipeBook: atomic.Value(data.RecipeBook),

		selectedItemSlot: atomicInt32(data.SelectedItemSlot),

		health:         atomicFloat32(data.Health),
		food:           atomicInt32(data.FoodLevel),
		foodExhaustion: atomicFloat32(data.FoodExhaustionLevel),
		foodSaturation: atomicFloat32(data.FoodSaturationLevel),

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
	return u64f(p.x.Load()), u64f(p.y.Load()), u64f(p.z.Load())
}

func (p *Player) Rotation() (yaw, pitch float32) {
	return u32f(p.yaw.Load()), u32f(p.pitch.Load())
}

func (p *Player) OnGround() bool {
	return p.onGround.Load()
}

func (p *Player) SetPosition(x, y, z float64) {
	p.x.Store(f64u(x))
	p.y.Store(f64u(y))
	p.z.Store(f64u(z))
}

func (p *Player) SetRotation(yaw, pitch float32) {
	p.yaw.Store(f32u(yaw))
	p.pitch.Store(f32u(pitch))
}

func (p *Player) SetOnGround(val bool) {
	p.onGround.Store(val)
}

func (p *Player) Motion() (x, y, z float64) {
	return u64f(p.vx.Load()), u64f(p.vy.Load()), u64f(p.vz.Load())
}

func (p *Player) SetMotion(x, y, z float64) {
	p.vx.Store(f64u(x))
	p.vy.Store(f64u(y))
	p.vz.Store(f64u(z))
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
	return u32f(p.health.Load())
}

func (p *Player) SetHealth(h float32) {
	p.health.Store(f32u(h))
}

func (p *Player) Food() int32 {
	return p.food.Load()
}

func (p *Player) SetFood(f int32) {
	p.food.Store(f)
}

func (p *Player) FoodSaturation() float32 {
	return u32f(p.foodSaturation.Load())
}

func (p *Player) SetFoodSaturation(fs float32) {
	p.foodSaturation.Store(f32u(fs))
}

func (p *Player) FoodExhaustion() float32 {
	return u32f(p.foodExhaustion.Load())
}

func (p *Player) SetFoodExhaustion(fh float32) {
	p.foodExhaustion.Store(f32u(fh))
}

func (p *Player) Abilities() level.PlayerAbilities {
	return p.abilities.Get()
}

func (p *Player) SetAbilities(abs level.PlayerAbilities) {
	p.abilities.Set(abs)
}

func (p *Player) GameMode() level.GameMode {
	return level.GameMode(p.gameMode.Load())
}

func (p *Player) SetGameMode(mode level.GameMode) {
	p.gameMode.Store(int32(mode))
}

func (p *Player) Attribute(id string) *entity.Attribute {
	p.att_mu.Lock()
	defer p.att_mu.Unlock()
	i := slices.IndexFunc(p.attributes, func(att entity.Attribute) bool { return att.Id == id })
	if i == -1 {
		attr, ok := entity.DefaultAttributes[id]
		if !ok {
			return nil
		}
		a := entity.Attribute{
			Id:   id,
			Base: attr,
		}
		p.attributes = append(p.attributes, a)
		return &a
	}
	attr := &p.attributes[i]
	return attr
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
	slot := p.selectedItemSlot.Load()
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
	p.selectedItemSlot.Store(slot)
}

func (p *Player) sync() {
	x, y, z := p.Position()
	yaw, pitch := p.Rotation()
	vx, vy, vz := p.Motion()

	p.data.Abilities = p.abilities.Get()
	p.data.Pos = [3]float64{x, y, z}
	p.data.Rotation = [2]float32{yaw, pitch}
	p.data.OnGround = p.onGround.Load()
	p.data.Dimension = p.dimension.Get()
	p.data.Inventory = *p.inventory
	p.data.RecipeBook = p.recipeBook.Get()

	p.data.Motion = [3]float64{vx, vy, vz}

	p.att_mu.RLock()
	p.data.Attributes = p.attributes
	p.att_mu.RUnlock()

	p.data.Health = u32f(p.health.Load())
	p.data.FoodLevel = p.food.Load()
	p.data.FoodExhaustionLevel = u32f(p.foodExhaustion.Load())
	p.data.FoodSaturationLevel = u32f(p.foodSaturation.Load())
	p.data.PlayerGameType = level.GameMode(p.gameMode.Load())
	p.data.SelectedItemSlot = p.selectedItemSlot.Load()
}
