package world

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"os"

	"github.com/aimjel/minecraft/nbt"
	"github.com/google/uuid"
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
	path                  string             `nbt:"-"`
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
	buf := bytes.NewBuffer(nil)
	//writer := gzip.NewWriter(buf)
	enc := nbt.NewEncoder(buf)
	fmt.Println(enc.Encode(*data))
	os.WriteFile(data.path, buf.Bytes(), 0755)
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

func UUIDToNBT(str string) (u []int32) {
	uuid, _ := uuid.Parse(str)
	s1 := uuid[:4]
	s2 := uuid[4:8]
	s3 := uuid[8:12]
	s4 := uuid[12:]
	u = append(u, int32(binary.BigEndian.Uint32(s1)))
	u = append(u, int32(binary.BigEndian.Uint32(s2)))
	u = append(u, int32(binary.BigEndian.Uint32(s3)))
	u = append(u, int32(binary.BigEndian.Uint32(s4)))
	return
}

func NBTToUUID(u []int32) (uuid.UUID, error) {
	byteSlice := make([]byte, 16)
	binary.BigEndian.PutUint32(byteSlice[:4], uint32(u[0]))
	binary.BigEndian.PutUint32(byteSlice[4:8], uint32(u[1]))
	binary.BigEndian.PutUint32(byteSlice[8:12], uint32(u[2]))
	binary.BigEndian.PutUint32(byteSlice[12:], uint32(u[3]))

	uuid, err := uuid.FromBytes(byteSlice)
	if err != nil {
		return uuid, err
	}

	return uuid, nil
}

func (world *World) GeneratePlayerData(uuid string) *PlayerData {
	return &PlayerData{
		path: fmt.Sprintf("world/playerdata/%s.dat", uuid),
		Pos: []float64{
			float64(world.nbt.Data.SpawnX),
			float64(world.nbt.Data.SpawnY),
			float64(world.nbt.Data.SpawnZ),
		},
		Rotation: []float32{
			0,
			0,
		},
		Motion: []float64{
			0,
			0,
			0,
		},
		Health:              20,
		Fire:                -20,
		FoodLevel:           20,
		FoodSaturationLevel: 5,
		OnGround:            1,
		UUID:                UUIDToNBT(uuid),
		PlayerGameType:      int32(world.Gamemode),
		Dimension:           "minecraft:overworld",
	}
}
