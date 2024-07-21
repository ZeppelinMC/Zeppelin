package player

import (
	"sync"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/world"
)

var _ entity.LivingEntity = (*Player)(nil)

type Player struct {
	entityId int32

	data       world.PlayerData
	x, y, z    atomic.AtomicValue[float64]
	yaw, pitch atomic.AtomicValue[float32]

	health         atomic.AtomicValue[float32]
	food           atomic.AtomicValue[int32]
	foodExhaustion atomic.AtomicValue[float32]
	foodSaturation atomic.AtomicValue[float32]

	clientInfo atomic.AtomicValue[configuration.ClientInformation]

	dimension atomic.AtomicValue[string]

	md_mu    sync.Mutex
	metadata metadata.Metadata
}

// NewPlayer creates a player struct with the entity id specified and initalizes an entity metadata map for it
func NewPlayer(entityId int32, data world.PlayerData) *Player {
	return &Player{
		entityId: entityId,
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
			// Player extends Living Entity
			metadata.PlayerAdditionalHeartsIndex:   metadata.Float(0),
			metadata.PlayerScoreIndex:              metadata.VarInt(0),
			metadata.PlayerDisplayedSkinPartsIndex: metadata.Byte(0),
			metadata.PlayerMainHandIndex:           metadata.Byte(1),
		},

		x: atomic.Value(data.Pos[0]),
		y: atomic.Value(data.Pos[1]),
		z: atomic.Value(data.Pos[2]),

		yaw:   atomic.Value(data.Rotation[0]),
		pitch: atomic.Value(data.Rotation[1]),

		dimension: atomic.Value(data.Dimension),

		health:         atomic.Value(data.Health),
		food:           atomic.Value(data.FoodLevel),
		foodExhaustion: atomic.Value(data.FoodExhaustionLevel),
		foodSaturation: atomic.Value(data.FoodSaturationLevel),
	}
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

func (p *Player) SetPosition(x, y, z float64) {
	p.x.Set(x)
	p.y.Set(y)
	p.z.Set(z)
}

func (p *Player) SetRotation(yaw, pitch float32) {
	p.yaw.Set(yaw)
	p.pitch.Set(pitch)
}

func (p *Player) EntityId() int32 {
	return p.entityId
}

func (p *Player) SetClientInformation(info configuration.ClientInformation) {
	p.clientInfo.Set(info)
}

func (p *Player) ClientInformation() configuration.ClientInformation {
	return p.clientInfo.Get()
}

func (p *Player) Metadata() metadata.Metadata {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	return p.metadata
}

func (p *Player) SetMetadata(md metadata.Metadata) {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
	p.metadata = md
}

func (p *Player) MetadataIndex(i byte) any {
	p.md_mu.Lock()
	defer p.md_mu.Unlock()
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
