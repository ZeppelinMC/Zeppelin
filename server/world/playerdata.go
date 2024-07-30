package world

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/entity"
)

func NewDataUUID(u uuid.UUID) DataUUID {
	return DataUUID{
		int32(u[0])<<24 | int32(u[1])<<16 | int32(u[2])<<8 | int32(u[3]),
		int32(u[4])<<24 | int32(u[5])<<16 | int32(u[6])<<8 | int32(u[7]),
		int32(u[8])<<24 | int32(u[9])<<16 | int32(u[10])<<8 | int32(u[11]),
		int32(u[12])<<24 | int32(u[13])<<16 | int32(u[14])<<8 | int32(u[15]),
	}
}

type DataUUID [4]int32

func (u DataUUID) UUID() uuid.UUID {
	return uuid.UUID{
		byte(u[0] >> 24),
		byte(u[0] >> 16),
		byte(u[0] >> 8),
		byte(u[0]),

		byte(u[1] >> 24),
		byte(u[1] >> 16),
		byte(u[1] >> 8),
		byte(u[1]),

		byte(u[2] >> 24),
		byte(u[2] >> 16),
		byte(u[2] >> 8),
		byte(u[2]),

		byte(u[3] >> 24),
		byte(u[3] >> 16),
		byte(u[3] >> 8),
		byte(u[3]),
	}
}

type PlayerData struct {
	// the path of this playerdata file, not a field in the nbt
	path string `nbt:"-"`

	AbsorptionAmount float32
	Air              int16
	Brain            struct {
		Memories struct{} `nbt:"memories"`
	}
	DataVersion       int32
	DeathTime         int16
	Dimension         string
	EnderItems        container.Container
	FallFlying        bool
	Fire              int16
	Health            float32
	HurtByTimestamp   int32
	HurtTime          int16
	Inventory         container.Container
	Invulnerable      bool
	LastDeathLocation struct {
		Dimension string   `nbt:"dimension"`
		Pos       [3]int32 `nbt:"pos"`
	}
	Motion           [3]float64
	OnGround         bool
	PortalCooldown   int32
	Pos              [3]float64
	Rotation         [2]float32
	Score            int32
	SelectedItemSlot int32
	SleepTimer       int16
	UUID             DataUUID

	XpLevel int32
	XpP     float32
	XpSeed  int32
	XpTotal int32

	Abilities PlayerAbilities `nbt:"abilities"`

	ActiveEffects []struct {
		Duration      int32  `nbt:"duration"`
		Id            string `nbt:"id"`
		ShowIcon      bool   `nbt:"show_icon"`
		ShowParticles bool   `nbt:"show_particles"`
	} `nbt:"active_effects"`

	Attributes []entity.Attribute `nbt:"attributes"`

	CurrentImpulseContextResetGraceTime int32 `nbt:"current_impulse_context_reset_grace_time"`

	FoodExhaustionLevel float32 `nbt:"foodExhaustionLevel"`
	FoodLevel           int32   `nbt:"foodLevel"`
	FoodSaturationLevel float32 `nbt:"foodSaturationLevel"`
	FoodTickTimer       int32   `nbt:"foodTickTimer"`

	IgnoreFallDamageFromCurrentExplosion bool     `nbt:"ignore_fall_damage_from_current_explosion"`
	PlayerGameType                       GameType `nbt:"playerGameType"`

	RecipeBook RecipeBook `nbt:"recipeBook"`

	SeenCredits               bool `nbt:"seenCredits"`
	SpawnExtraParticlesOnFall bool `nbt:"spawn_extra_particles_on_fall"`
	WardenSpawnTracker        struct {
		CooldownTicks         int32 `nbt:"cooldown_ticks"`
		TicksSinceLastWarning int32 `nbt:"ticks_since_last_warning"`
		WarningLevel          int32 `nbt:"warning_level"`
	} `nbt:"warden_spawn_tracker"`
}

type RecipeBook struct {
	IsBlastingFurnaceFilteringCraftable bool `nbt:"isBlastingFurnaceFilteringCraftable"`
	IsBlastingFurnaceGuiOpen            bool `nbt:"isBlastingFurnaceGuiOpen"`
	IsFilteringCraftable                bool `nbt:"isFilteringCraftable"`
	IsFurnaceFilteringCraftable         bool `nbt:"isFurnaceFilteringCraftable"`
	IsFurnaceGuiOpen                    bool `nbt:"isFurnaceGuiOpen"`
	IsGuiOpen                           bool `nbt:"isGuiOpen"`
	IsSmokerFilteringCraftable          bool `nbt:"isSmokerFilteringCraftable"`
	IsSmokerGuiOpen                     bool `nbt:"isSmokerGuiOpen"`

	Recipes       []string `nbt:"recipes"`
	ToBeDisplayed []string `nbt:"toBeDisplayed"`
}

func (data *PlayerData) Save() error {
	file, err := os.Create(data.path)
	if err != nil {
		return err
	}
	gzip := gzip.NewWriter(file)
	if err := nbt.NewEncoder(gzip).Encode("", *data); err != nil {
		return err
	}

	if err := gzip.Close(); err != nil {
		return err
	}
	return file.Close()
}

func (w *World) PlayerData(uuid string) (PlayerData, error) {
	var playerData PlayerData
	path := fmt.Sprintf("%s/playerdata/%s.dat", w.path, uuid)

	file, err := os.Open(path)
	if err != nil {
		return playerData, err
	}
	rd, err := gzip.NewReader(file)
	if err != nil {
		return playerData, err
	}

	var buf, _ = io.ReadAll(rd)

	rd.Close()
	file.Close()

	_, err = nbt.NewDecoder(bytes.NewReader(buf)).Decode(&playerData)
	playerData.path = path

	return playerData, err
}

func (w *World) NewPlayerData(uuid uuid.UUID) PlayerData {
	return PlayerData{
		path: fmt.Sprintf("%s/playerdata/%s.dat", w.path, uuid),
		Pos:  [3]float64{float64(w.Data.SpawnX), float64(w.Data.SpawnY), float64(w.Data.SpawnZ)},

		Health:              20,
		FoodSaturationLevel: 5,
		FoodLevel:           20,
		Fire:                -20,

		UUID:           NewDataUUID(uuid),
		Dimension:      "minecraft:overworld",
		OnGround:       true,
		PlayerGameType: w.Data.GameType,
		Abilities: PlayerAbilities{
			FlySpeed:     0.05,
			Instabuild:   w.Data.GameType == GameTypeCreative,
			Invulnerable: w.Data.GameType == GameTypeCreative,
			MayFly:       w.Data.GameType == GameTypeCreative,
			MayBuild:     w.Data.GameType != GameTypeAdventure,
			WalkSpeed:    0.1,
		},
		Attributes: []entity.Attribute{
			{
				Base: 4.5,
				Id:   "minecraft:player.block_interaction_range",
			},
			{
				Base: 0.1,
				Id:   "minecraft:generic.movement_speed",
			},
			{
				Base: 3,
				Id:   "minecraft:player.entity_interaction_range",
			},
		},
	}
}

type PlayerAbilities struct {
	FlySpeed     float32 `nbt:"flySpeed"`
	Flying       bool    `nbt:"flying"`
	Instabuild   bool    `nbt:"instabuild"`
	Invulnerable bool    `nbt:"invulnerable"`
	MayBuild     bool    `nbt:"mayBuild"`
	MayFly       bool    `nbt:"mayfly"`
	WalkSpeed    float32 `nbt:"walkSpeed"`
}

func (a PlayerAbilities) Encode(fovModifier float32) *play.PlayerAbilitiesClientbound {
	var flags int8
	if a.Invulnerable {
		flags |= 0x01
	}
	if a.Flying {
		flags |= 0x02
	}
	if a.MayFly {
		flags |= 0x04
	}
	if a.Instabuild {
		flags |= 0x08
	}

	return &play.PlayerAbilitiesClientbound{
		Flags:       flags,
		FlyingSpeed: a.FlySpeed,
		FOVModifier: fovModifier,
	}
}
