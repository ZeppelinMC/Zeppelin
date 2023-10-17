package world

import (
	"bytes"
	"compress/gzip"
	"math"
	"math/rand"
	"os"

	_ "embed"

	"github.com/aimjel/minecraft/nbt"
	"github.com/dynamitemc/dynamite/server/world/chunk"
)

func booltoint(b bool) int {
	if b {
		return 1
	}
	return 0
}

func GenerateWorldData(hardcore int8) worldData {
	return worldData{
		Data: WorldData{
			ServerBrands: []string{"Dynamite"},
			//SpawnY:              72,
			//ThunderTime:         19779,
			//LastPlayed:          1693159101540,
			//BorderSize:          5.9999968e+07,
			DataVersion: 3465,
			//Time:                938,
			Difficulty:          2,
			Raining:             0,
			BorderWarningBlocks: 5,
			WorldGenSettings: WorldGenSettings{
				BonusChest:       0,
				Seed:             int64(rand.Float64() * math.MaxFloat64),
				GenerateFeatures: 1,
				Dimensions: Dimensions{
					End: DimensionData{
						Generator: DimensionGenerator{
							BiomeSource: BiomeSource{
								Preset: "",
								Type:   "minecraft:the_end",
							},
							Type: "minecraft:noise",
							//Settings: "minecraft:end",
						}, Type: "minecraft:the_end",
					},
					Overworld: DimensionData{
						Generator: DimensionGenerator{
							BiomeSource: BiomeSource{
								Preset: "minecraft:overworld",
								Type:   "minecraft:multi_noise",
							},
							Type: "minecraft:noise",
							//Settings: "minecraft:overworld",
						}, Type: "minecraft:overworld"},
					Nether: DimensionData{
						Generator: DimensionGenerator{
							BiomeSource: BiomeSource{
								Preset: "minecraft:nether",
								Type:   "minecraft:multi_noise",
							},
							Type: "minecraft:noise",
							//Settings: "minecraft:nether",
						},
						Type: "minecraft:the_nether",
					},
				},
			},
			//BorderSizeLerpTarget: 5.9999968e+07,
			Version: Version{
				Snapshot: 0,
				Series:   "main",
				Id:       3465,
				Name:     "1.20.1",
			},
			//DayTime:                    938,
			WanderingTraderSpawnChance: 25,
			Hardcore:                   hardcore,
			//SpawnX:                     176,
			BorderWarningTime:         15,
			WanderingTraderSpawnDelay: 24000,
			CustomBossEvents:          struct{}{},
			ClearWeatherTime:          0,
			Thundering:                0,
			SpawnAngle:                0,
			ScheduledEvents:           nil,
			BorderCenterZ:             0,
			//RainTime:                   102381,
			LevelName:     "world",
			WasModded:     0,
			BorderCenterX: 0,
			AllowCommands: 1,
			DragonFight: DragonFight{
				NeedsStateScanning: 1,
				//Gateways:           []int32{14, 17, 11, 1, 10, 6, 9, 4, 16, 18, 12, 13, 5, 19, 0, 8, 15, 3, 7, 2},
				DragonKilled:     0,
				PreviouslyKilled: 0,
			},
			BorderDamagePerBlock: 0.2,
			Initialized:          1,
			GameRules: map[string]string{
				"doDaylightCycle":               "true",
				"doInsomnia":                    "true",
				"doEntityDrops":                 "true",
				"spectatorsGenerateChunks":      "true",
				"disableElytraMovementCheck":    "false",
				"disableRaids":                  "false",
				"commandModificationBlockLimit": "32768",
				"forgiveDeadPlayers":            "true",
				"maxEntityCramming":             "24",
				"spawnRadius":                   "10",
				"announceAdvancements":          "true",
				"universalAnger":                "false",
				"fallDamage":                    "true",
				"randomTickSpeed":               "3",
				"doTraderSpawning":              "true",
				"drowningDamage":                "true",
				"showDeathMessages":             "true",
				"playersSleepingPercentage":     "100",
				"doImmediateRespawn":            "false",
				"naturalRegeneration":           "true",
				"keepInventory":                 "false",
				"doVinesSpread":                 "true",
				"doTileDrops":                   "true",
				"maxCommandChainLength":         "65536",
				"fireDamage":                    "true",
				"waterSourceConversion":         "true",
				"commandBlockOutput":            "true",
				"doWeatherCycle":                "true",
				"snowAccumulationHeight":        "1",
				"doWardenSpawning":              "true",
				"doFireTick":                    "true",
				"doPatrolSpawning":              "true",
				"doMobLoot":                     "true",
				"sendCommandFeedback":           "true",
				"doMobSpawning":                 "true",
				"mobExplosionDropDecay":         "true",
				"mobGriefing":                   "true",
				"freezeDamage":                  "true",
				"logAdminCommands":              "true",
				"globalSoundEvents":             "true",
				"tntExplosionDropDecay":         "false",
				"blockExplosionDropDecay":       "true",
				"doLimitedCrafting":             "false",
				"reducedDebugInfo":              "false",
				"lavaSourceConversion":          "false",
			},
			//SpawnZ:         208,
			BorderSafeZone: 5,
			DataPacks: DataPacks{
				Disabled: []string{"bundle"},
				Enabled:  []string{"vanilla"},
			},
			GameType:           1,
			DifficultyLocked:   0,
			BorderSizeLerpTime: 0,
			VersionNumber:      19133,
		},
	}
}

func CreateWorld(hardcore bool) {
	hc := int8(booltoint(hardcore))
	world := GenerateWorldData(hc)

	os.Mkdir("world", 0755)
	os.Mkdir("world/data", 0755)
	os.Mkdir("world/datapacks", 0755)
	os.Mkdir("world/entities", 0755)
	os.Mkdir("world/playerdata", 0755)
	os.Mkdir("world/poi", 0755)
	os.Mkdir("world/region", 0755)

	os.WriteFile("world/session.lock", nil, 0755)

	file, _ := os.Create("world/level.dat")

	buf := bytes.NewBuffer(nil)
	enc := nbt.NewEncoder(buf)
	enc.Encode(world)

	writer := gzip.NewWriter(file)
	writer.Write(buf.Bytes())
	writer.Close()
	file.Close()
}

func (w *World) Save() {
	buf := bytes.NewBuffer(nil)
	writer := gzip.NewWriter(buf)
	enc := nbt.NewEncoder(writer)
	enc.Encode(w.nbt)
	os.WriteFile("world/level.dat", buf.Bytes(), 0755)
}

type Generator interface {
	Generate(x, z int32) (*chunk.Chunk, error)
}

type FlatGenerator struct{}

//go:embed flatchunk.nbt
var flatchunk []byte

func (f *FlatGenerator) Generate(x, z int32) (*chunk.Chunk, error) {
	return chunk.NewAnvilChunk(flatchunk)
}
