package player

import (
	"fmt"
	"math"
	"sync"

	"github.com/dynamitemc/dynamite/server/world"
)

type Player struct {
	isHardCore bool
	gameMode   byte

	dead           bool
	health         float32
	food           int32
	foodSaturation float32

	data *world.PlayerData

	inventory []world.Slot

	dimension *world.Dimension

	viewDistance       int32
	simulationDistance int32

	entityID int32

	operator bool

	clientSettings ClientInformation

	x, y, z    float64
	yaw, pitch float32
	onGround   bool

	flying bool

	mu sync.RWMutex
}

type ClientInformation struct {
	Locale               string
	ViewDistance         int8
	ChatMode             int32
	ChatColors           bool
	DisplayedSkinParts   uint8
	MainHand             int32
	DisableTextFiltering bool
	AllowServerListings  bool
}

func New(entityID int32, vd, sd int32, data *world.PlayerData) *Player {
	pl := &Player{entityID: entityID, viewDistance: vd, simulationDistance: sd, data: data}
	pl.inventory = data.Inventory
	pl.gameMode = byte(data.PlayerGameType)
	pl.x, pl.y, pl.z, pl.yaw, pl.pitch = data.Pos[0], data.Pos[1], data.Pos[2], data.Rotation[0], data.Rotation[1]
	pl.health, pl.food, pl.foodSaturation = data.Health, data.FoodLevel, data.FoodSaturationLevel
	fl := true
	if data.Abilities.Flying == 0 {
		fl = false
	}
	pl.flying = fl
	return pl
}

func (p *Player) Dimension() *world.Dimension {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.dimension
}

func (p *Player) IsDead() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.dead
}

func (p *Player) SetDead(a bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dead = a
}

func (p *Player) SetDimension(d *world.Dimension) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.dimension = d
}

func (p *Player) Inventory() []world.Slot {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.inventory
}

func (p *Player) SetInventory(i []world.Slot) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.inventory = i
}

func (p *Player) ClientSettings() ClientInformation {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.clientSettings
}

func (p *Player) SetClientSettings(information ClientInformation) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clientSettings = information
}

func (p *Player) Health() float32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.health
}

func (p *Player) SetHealth(health float32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.health = health
}

func (p *Player) FoodLevel() int32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.food
}

func (p *Player) SetFoodLevel(level int32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.food = level
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

func (p *Player) ViewDistance() int32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.viewDistance
}

func (p *Player) SavedPosition() (x, y, z float64) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.data.Pos[0], p.data.Pos[1], p.data.Pos[2]
}

func (p *Player) SavedRotation() (yaw, pitch float32) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.data.Rotation[0], p.data.Rotation[1]
}

func (p *Player) SavedOnGround() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.data.OnGround != 0
}

func (p *Player) SavedAbilities() world.Abilities {
	return p.data.Abilities
}

func (p *Player) SetFlying(val bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.flying = val
}

func (p *Player) Save() {
	o := int8(0)
	if p.onGround {
		o = 1
	}
	fl := int8(0)
	if p.flying {
		fl = 1
	}

	p.data.Pos[0], p.data.Pos[1], p.data.Pos[2], p.data.Rotation[0], p.data.Rotation[1], p.data.OnGround = p.x, p.y, p.z, p.yaw, p.pitch, o
	p.data.PlayerGameType = int32(p.gameMode)
	p.data.Inventory = p.inventory
	p.data.Abilities.Flying = fl

	p.data.Save()
}

func (p *Player) SimulationDistance() int32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.simulationDistance
}

func (p *Player) IsHardcore() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.isHardCore
}

func (p *Player) SetGameMode(gm byte) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.gameMode = gm
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
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.onGround
}

func (p *Player) SetPosition(x, y, z float64, yaw, pitch float32, ong bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.x, p.y, p.z, p.yaw, p.pitch, p.onGround = x, y, z, yaw, pitch, ong
}

func (p *Player) GetPosition2() uint64 {
	x := int64(math.Float64bits(p.x))
	y := int64(math.Float64bits(p.y))
	z := int64(math.Float64bits(p.z))
	fmt.Println(x, y, z)
	return uint64((x << 38) | (z << 12) | (y))
}

func (p *Player) Operator() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.operator
}

func (p *Player) SetOperator(op bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.operator = op
}

func (p *Player) EntityId() int32 {
	//no need to protect this with mutex because it never changes
	return p.entityID
}
