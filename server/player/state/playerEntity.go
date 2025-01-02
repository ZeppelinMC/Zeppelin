package state

import (
	a "sync/atomic"

	"github.com/zeppelinmc/zeppelin/protocol/net/metadata"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/registry"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/util/atomic"
)

type PlayerEntity struct {
	entity.LivingEntity

	data level.Player

	abilities atomic.AtomicValue[level.PlayerAbilities]

	gameMode         a.Int32
	selectedItemSlot a.Int32

	//recipeBook atomic.AtomicValue[level.RecipeBook]

	inventory *container.Container
}

// New looks up a player state in the cache or creates one if not found
func (mgr *PlayerEntityManager) New(data level.Player) *PlayerEntity {
	if p, ok := mgr.lookup(data.UUID.UUID()); ok {
		return p
	}

	pl := &PlayerEntity{
		LivingEntity: entity.NewLiving(entity.New(data.UUID.UUID(), registry.EntityType.Get("minecraft:player"), metadata.Player(data.Health), data.Attributes)),

		//recipeBook: atomic.Value(data.RecipeBook),

		abilities: atomic.Value(data.Abilities),

		inventory: &data.Inventory,

		data: data,
	}

	pl.gameMode.Store(int32(data.PlayerGameType))
	pl.selectedItemSlot.Store(data.SelectedItemSlot)

	pl.SetPosition(data.Pos[0], data.Pos[1], data.Pos[2])
	pl.SetMotion(data.Motion[0], data.Motion[1], data.Motion[2])
	pl.SetRotation(data.Rotation[0], data.Rotation[1])
	pl.SetOnGround(data.OnGround)
	pl.SetDimensionName(data.Dimension)
	pl.SetHealth(data.Health)
	pl.SetFood(data.FoodLevel, data.FoodSaturationLevel, data.FoodExhaustionLevel)

	mgr.add(pl)

	return pl
}

func (p *PlayerEntity) Abilities() level.PlayerAbilities {
	return p.abilities.Get()
}

func (p *PlayerEntity) SetAbilities(abs level.PlayerAbilities) {
	p.abilities.Set(abs)
}

func (p *PlayerEntity) GameMode() level.GameMode {
	return level.GameMode(p.gameMode.Load())
}

func (p *PlayerEntity) SetGameMode(mode level.GameMode) {
	p.gameMode.Store(int32(mode))
}

/*func (p *PlayerEntity) RecipeBook() level.RecipeBook {
	return p.recipeBook.Get()
}

func (p *PlayerEntity) SetRecipeBook(book level.RecipeBook) {
	p.recipeBook.Set(book)
}*/

func (p *PlayerEntity) Inventory() *container.Container {
	return p.inventory
}

// if negative, returns 0 and if over 8, returns 8
func (p *PlayerEntity) SelectedItemSlot() int32 {
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
func (p *PlayerEntity) SetSelectedItemSlot(slot int32) {
	if slot < 0 {
		slot = 0
	}
	if slot > 8 {
		slot = 8
	}
	p.selectedItemSlot.Store(slot)
}

func (p *PlayerEntity) sync() {
	x, y, z := p.Position()
	yaw, pitch := p.Rotation()
	vx, vy, vz := p.Motion()

	p.data.Abilities = p.abilities.Get()
	p.data.Pos = [3]float64{x, y, z}
	p.data.Rotation = [2]float32{yaw, pitch}
	p.data.OnGround = p.OnGround()
	p.data.Dimension = p.DimensionName()
	p.data.Inventory = *p.inventory
	//p.data.RecipeBook = p.recipeBook.Get()

	p.data.Motion = [3]float64{vx, vy, vz}

	p.data.Attributes = p.Attributes()

	p.data.Health = p.Health()
	p.data.FoodLevel, p.data.FoodSaturationLevel, p.data.FoodExhaustionLevel = p.Food()
	p.data.PlayerGameType = level.GameMode(p.gameMode.Load())
	p.data.SelectedItemSlot = p.selectedItemSlot.Load()
}
