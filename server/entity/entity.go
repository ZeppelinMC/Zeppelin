package entity

import (
	"maps"
	"math"
	"slices"
	"sync"
	"sync/atomic"

	a "github.com/zeppelinmc/zeppelin/util/atomic"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
)

var entityId atomic.Int32

func NewEntityId() int32 {
	return entityId.Add(1)
}

func New(uuid uuid.UUID, typ int32, metadata metadata.Metadata, attributes []Attribute) *Entity {
	return &Entity{
		uuid:       uuid,
		typ:        typ,
		metadata:   metadata,
		attributes: attributes,
	}
}

func NewLiving(e *Entity) LivingEntity {
	return LivingEntity{Entity: *e}
}

// Entity is the base state for entities
type Entity struct {
	uuid uuid.UUID
	typ  int32

	x, y, z    atomic.Uint64 // the position of the state
	vx, vy, vz atomic.Uint64 // the velocity of the state
	yaw, pitch atomic.Uint32 // the rotation of the state
	onGround   atomic.Bool

	md_mu    sync.RWMutex
	metadata metadata.Metadata

	attr_mu    sync.RWMutex
	attributes []Attribute

	dimensionName a.AtomicValue[string]
}

func (p *Entity) Attribute(id string) *Attribute {
	p.attr_mu.Lock()
	defer p.attr_mu.Unlock()
	i := slices.IndexFunc(p.attributes, func(att Attribute) bool { return att.Id == id })
	if i == -1 {
		attr, ok := DefaultAttributes[id]
		if !ok {
			return nil
		}
		a := Attribute{
			Id:   id,
			Base: attr,
		}
		p.attributes = append(p.attributes, a)
		return &a
	}
	attr := &p.attributes[i]
	return attr
}

// Attributes returns a clone of the attributes of this state
func (p *Entity) Attributes() []Attribute {
	p.attr_mu.RLock()
	defer p.attr_mu.RUnlock()
	return slices.Clone(p.attributes)
}

func (p *Entity) SetAttribute(id string, base float64) {
	p.attr_mu.Lock()
	defer p.attr_mu.Unlock()
	i := slices.IndexFunc(p.attributes, func(att Attribute) bool { return att.Id == id })
	if i == -1 {
		return
	}
	p.attributes[i].Base = base
}

// Metadata returns a clone of the metadata of this state
func (p *Entity) Metadata() metadata.Metadata {
	p.md_mu.RLock()
	defer p.md_mu.RUnlock()
	return maps.Clone(p.metadata)
}

func (p *Entity) SetMetadata(md metadata.Metadata) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	p.metadata = md
}

func (p *Entity) MetadataIndex(i byte) any {
	p.md_mu.RLock()
	defer p.md_mu.RUnlock()
	return p.metadata[i]
}

func (p *Entity) SetMetadataIndex(i byte, v any) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	p.metadata[i] = v
}

func (p *Entity) SetMetadataIndexes(md metadata.Metadata) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	for index, value := range md {
		p.metadata[index] = value
	}
}

func (p *Entity) DimensionName() string {
	return p.dimensionName.Get()
}

func (p *Entity) SetDimensionName(v string) {
	p.dimensionName.Set(v)
}

func (p *Entity) SetMotion(x, y, z float64) {
	p.vx.Store(math.Float64bits(x))
	p.vy.Store(math.Float64bits(y))
	p.vz.Store(math.Float64bits(z))
}

func (p *Entity) SetPosition(x, y, z float64) {
	p.x.Store(math.Float64bits(x))
	p.y.Store(math.Float64bits(y))
	p.z.Store(math.Float64bits(z))
}

func (p *Entity) SetRotation(yaw, pitch float32) {
	p.yaw.Store(math.Float32bits(yaw))
	p.pitch.Store(math.Float32bits(pitch))
}

func (p *Entity) SetOnGround(v bool) {
	p.onGround.Store(v)
}

func (p *Entity) Position() (x, y, z float64) {
	return math.Float64frombits(p.x.Load()), math.Float64frombits(p.y.Load()), math.Float64frombits(p.z.Load())
}

func (p *Entity) Motion() (vx, vy, vz float64) {
	return math.Float64frombits(p.vx.Load()), math.Float64frombits(p.vy.Load()), math.Float64frombits(p.vz.Load())
}

func (p *Entity) Rotation() (yaw, pitch float32) {
	return math.Float32frombits(p.yaw.Load()), math.Float32frombits(p.pitch.Load())
}

func (p *Entity) OnGround() bool {
	return p.onGround.Load()
}

func (p *Entity) UUID() uuid.UUID {
	return p.uuid
}

func (p *Entity) Type() int32 {
	return p.typ
}

type LivingEntity struct {
	Entity
	health,
	food,
	foodSaturation, foodExhaustion atomic.Uint32
}

func (l *LivingEntity) SetHealth(h float32) {
	l.health.Store(math.Float32bits(h))
}

func (l *LivingEntity) SetFood(food int32, saturation, exhaustion float32) {
	l.food.Store(uint32(food))
	l.foodSaturation.Store(math.Float32bits(saturation))
	l.foodExhaustion.Store(math.Float32bits(exhaustion))
}

func (l *LivingEntity) Health() float32 {
	return math.Float32frombits(l.health.Load())
}

func (l *LivingEntity) Food() (food int32, saturation, exhaustion float32) {
	return int32(l.food.Load()), math.Float32frombits(l.foodSaturation.Load()), math.Float32frombits(l.foodExhaustion.Load())
}

type Attribute struct {
	Base float64 `nbt:"base"`
	Id   string  `nbt:"id"`
}

var DefaultAttributes = map[string]float64{
	"minecraft:generic.armor":                     0,
	"minecraft:generic.armor_toughness":           0,
	"minecraft:generic.attack_damage":             2,
	"minecraft:generic.attack_knockback":          0,
	"minecraft:generic.attack_speed":              2,
	"minecraft:generic.block_break_speed":         1,
	"minecraft:generic.block_interaction_range":   4.5,
	"minecraft:generic.entity_interaction_range":  3,
	"minecraft:generic.fall_damage_multiplier":    1,
	"minecraft:generic.flying_speed":              0.4,
	"minecraft:generic.follow_range":              32,
	"minecraft:generic.gravity":                   0.08,
	"minecraft:generic.jump_strength":             0.42,
	"minecraft:generic.knockback_resistance":      0,
	"minecraft:generic.luck":                      0,
	"minecraft:generic.max_absorption":            0,
	"minecraft:generic.max_health":                20,
	"minecraft:generic.movement_speed":            0.7,
	"minecraft:generic.safe_fall_distance":        3,
	"minecraft:generic.scale":                     1,
	"minecraft:zombie.spawn_reinforcements":       0,
	"minecraft:generic.step_height":               0.6,
	"minecraft:generic.submerged_mining_speed":    0.2,
	"minecraft:generic.sweeping_damage_ratio":     0,
	"minecraft:generic.water_movement_efficiency": 0,
}
