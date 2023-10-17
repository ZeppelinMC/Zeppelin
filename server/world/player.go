package world

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/aimjel/minecraft/nbt"
	uuid2 "github.com/google/uuid"
)

type Slot struct {
	Count int8    `nbt:"Count"`
	Slot  int8    `nbt:"Slot"`
	Id    string  `nbt:"id"`
	Tag   SlotTag `nbt:"tag"`
}

type Enchantment struct {
	Id    string `nbt:"id"`
	Level int32  `nbt:"lvl"`
}

type SlotTag struct {
	Damage       int32         `nbt:"Damage"`
	RepairCost   int32         `nbt:"RepairCost"`
	Enchantments []Enchantment `nbt:"Enchantment"`
}

type RecipeBook struct {
	IsGuiOpen                           int8     `nbt:"isGuiOpen"`
	IsBlastingFurnaceGuiOpen            int8     `nbt:"isBlastingFurnaceGuiOpen"`
	IsSmokerGuiOpen                     int8     `nbt:"isSmokerGuiOpen"`
	IsBlastingFurnaceFilteringCraftable int8     `nbt:"isBlastingFurnaceFilteringCraftable"`
	IsFilteringCraftable                int8     `nbt:"isFilteringCraftable"`
	ToBeDisplayed                       []string `nbt:"toBeDisplayed"`
	IsFurnaceGuiOpen                    int8     `nbt:"isFurnaceGuiOpen"`
	IsFurnaceFilteringCraftable         int8     `nbt:"isFurnaceFilteringCraftable"`
	IsSmokerFilteringCraftable          int8     `nbt:"isSmokerFilteringCraftable"`
	Recipes                             []string `nbt:"recipes"`
}

type WardenSpawnTracker struct {
	WarningLevel          int32 `nbt:"warning_level"`
	TicksSinceLastWarning int32 `nbt:"ticks_since_last_warning"`
	CooldownTicks         int32 `nbt:"cooldown_ticks"`
}

type Attribute struct {
	Name string  `nbt:"Name"`
	Base float64 `nbt:"Base"`
}

type Abilities struct {
	Mayfly       int8    `nbt:"mayfly"`
	Instabuild   int8    `nbt:"instabuild"`
	WalkSpeed    float32 `nbt:"walkSpeed"`
	MayBuild     int8    `nbt:"mayBuild"`
	Flying       int8    `nbt:"flying"`
	FlySpeed     float32 `nbt:"flySpeed"`
	Invulnerable int8    `nbt:"invulnerable"`
}

type Brain struct {
	Memories struct{} `nbt:"memories"`
}

type PlayerData struct {
	path                  string
	Invulnerable          int8               `nbt:"Invulnerable"`
	FoodSaturationLevel   float32            `nbt:"foodSaturationLevel"`
	UUID                  []int32            `nbt:"UUID"`
	EnderItems            []Slot             `nbt:"EnderItems"`
	DataVersion           int32              `nbt:"DataVersion"`
	SelectedItemSlot      int32              `nbt:"SelectedItemSlot"`
	SleepTimer            int16              `nbt:"SleepTimer"`
	Abilities             Abilities          `nbt:"abilities"`
	RecipeBook            RecipeBook         `nbt:"recipeBook"`
	XpSeed                int32              `nbt:"XpSeed"`
	Inventory             []Slot             `nbt:"Inventory"`
	FoodLevel             int32              `nbt:"foodLevel"`
	HurtByTimestamp       int32              `nbt:"HurtByTimestamp"`
	FallDistance          float32            `nbt:"FallDistance"`
	PlayerGameType        int32              `nbt:"playerGameType"`
	SeenCredits           int8               `nbt:"seenCredits"`
	Pos                   []float64          `nbt:"Pos"`
	FoodTickTimer         int32              `nbt:"foodTickTimer"`
	Brain                 Brain              `nbt:"Brain"`
	AbsorptionAmount      float32            `nbt:"AbsorptionAmount"`
	DeathTime             int16              `nbt:"DeathTime"`
	XpLevel               int32              `nbt:"XpLevel"`
	XpP                   float32            `nbt:"XpP"`
	FallFlying            int8               `nbt:"FallFlying"`
	Motion                []float64          `nbt:"Motion"`
	OnGround              int8               `nbt:"OnGround"`
	Rotation              []float32          `nbt:"Rotation"`
	Score                 int32              `nbt:"Score"`
	Fire                  int16              `nbt:"Fire"`
	FoodExhaustionLevel   float32            `nbt:"foodExhaustionLevel"`
	Attributes            []Attribute        `nbt:"Attributes"`
	EnteredNetherPosition []float64          `nbt:"enteredNetherPosition"`
	PortalCooldown        int32              `nbt:"PortalCooldown"`
	Health                float32            `nbt:"Health"`
	Dimension             string             `nbt:"Dimension"`
	XpTotal               int32              `nbt:"XpTotal"`
	Air                   int16              `nbt:"Air"`
	WardenSpawnTracker    WardenSpawnTracker `nbt:"warden_spawn_tracker"`
	HurtTime              int16              `nbt:"HurtTime"`
}

func (data *PlayerData) Save() {
	f, _ := os.Create(data.path)
	writer := gzip.NewWriter(f)
	buf := bytes.NewBuffer(nil)
	enc := nbt.NewEncoder(buf)
	enc.Encode(*data)
	writer.Write(buf.Bytes())

	writer.Close()
	f.Close()
}

func (world *World) GetPlayerData(uuid string) (data *PlayerData) {
	var d PlayerData
	f, err := os.Open(fmt.Sprintf("world/playerdata/%s.dat", uuid))
	if err != nil {
		data = world.GeneratePlayerData(uuid)
		return
	}
	gzipRd, err := gzip.NewReader(f)
	if err != nil {
		data = world.GeneratePlayerData(uuid)
		return
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(gzipRd); err != nil {
		data = world.GeneratePlayerData(uuid)
		return
	}

	err = nbt.Unmarshal(buf.Bytes(), &d)
	data = &d
	if err != nil {
		data = world.GeneratePlayerData(uuid)
	}
	data.path = fmt.Sprintf("world/playerdata/%s.dat", uuid)
	return
}

func ByteUUIDToIntUUID(uuid uuid2.UUID) (u []int32) {
	return append(u,
		int32(binary.BigEndian.Uint32(uuid[:4])),
		int32(binary.BigEndian.Uint32(uuid[4:8])),
		int32(binary.BigEndian.Uint32(uuid[8:12])),
		int32(binary.BigEndian.Uint32(uuid[12:])),
	)
}

func IntUUIDToByteUUID(u []int32) (uuid2.UUID, error) {
	byteSlice := make([]byte, 16)
	binary.BigEndian.PutUint32(byteSlice[:4], uint32(u[0]))
	binary.BigEndian.PutUint32(byteSlice[4:8], uint32(u[1]))
	binary.BigEndian.PutUint32(byteSlice[8:12], uint32(u[2]))
	binary.BigEndian.PutUint32(byteSlice[12:], uint32(u[3]))

	return uuid2.FromBytes(byteSlice)
}

func (world *World) GeneratePlayerData(uuid string) *PlayerData {
	u, _ := uuid2.Parse(uuid)
	return &PlayerData{
		path: fmt.Sprintf("world/playerdata/%s.dat", uuid),
		Pos: []float64{
			float64(world.nbt.Data.SpawnX),
			float64(world.nbt.Data.SpawnY),
			float64(world.nbt.Data.SpawnZ),
		},
		Rotation: []float32{
			90,
			0,
		},
		Health:              20,
		Fire:                -20,
		FoodLevel:           20,
		FoodSaturationLevel: 5,
		OnGround:            1,
		UUID:                ByteUUIDToIntUUID(u),
		PlayerGameType:      int32(world.Gamemode),
		Dimension:           "minecraft:overworld",
	}
}
