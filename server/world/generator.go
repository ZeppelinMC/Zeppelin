package world

import (
	"math"
	"math/rand"
	"os"
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
			//DataVersion:         3465,
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
							Type:     "minecraft:noise",
							Settings: "minecraft:end",
						}, Type: "minecraft:the_end",
					},
					Overworld: DimensionData{
						Generator: DimensionGenerator{
							BiomeSource: BiomeSource{
								Preset: "minecraft:overworld",
								Type:   "minecraft:multi_noise",
							},
							Type:     "minecraft:noise",
							Settings: "minecraft:overworld",
						}, Type: "minecraft:overworld"},
					Nether: DimensionData{
						Generator: DimensionGenerator{
							BiomeSource: BiomeSource{
								Preset: "minecraft:nether",
								Type:   "minecraft:multi_noise",
							},
							Type:     "minecraft:noise",
							Settings: "minecraft:nether",
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
			GameRules: GameRules{
				DoDaylightCycle:               "true",
				DoInsomnia:                    "true",
				DoEntityDrops:                 "true",
				SpectatorsGenerateChunks:      "true",
				DisableElytraMovementCheck:    "false",
				DisableRaids:                  "false",
				CommandModificationBlockLimit: "32768",
				ForgiveDeadPlayers:            "true",
				MaxEntityCramming:             "24",
				SpawnRadius:                   "10",
				AnnounceAdvancements:          "true",
				UniversalAnger:                "false",
				FallDamage:                    "true",
				RandomTickSpeed:               "3",
				DoTraderSpawning:              "true",
				DrowningDamage:                "true",
				ShowDeathMessages:             "true",
				PlayersSleepingPercentage:     "100",
				DoImmediateRespawn:            "false",
				NaturalRegeneration:           "true",
				KeepInventory:                 "false",
				DoVinesSpread:                 "true",
				DoTileDrops:                   "true",
				MaxCommandChainLength:         "65536",
				FireDamage:                    "true",
				WaterSourceConversion:         "true",
				CommandBlockOutput:            "true",
				DoWeatherCycle:                "true",
				SnowAccumulationHeight:        "1",
				DoWardenSpawning:              "true",
				DoFireTick:                    "true",
				DoPatrolSpawning:              "true",
				DoMobLoot:                     "true",
				SendCommandFeedback:           "true",
				DoMobSpawning:                 "true",
				MobExplosionDropDecay:         "true",
				MobGriefing:                   "true",
				FreezeDamage:                  "true",
				LogAdminCommands:              "true",
				GlobalSoundEvents:             "true",
				TntExplosionDropDecay:         "false",
				BlockExplosionDropDecay:       "true",
				DoLimitedCrafting:             "false",
				ReducedDebugInfo:              "false",
				LavaSourceConversion:          "false",
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
	_ = GenerateWorldData(hc)

	os.Mkdir("world", 0755)
	os.Mkdir("world/data", 0755)
	os.Mkdir("world/datapacks", 0755)
	os.Mkdir("world/entities", 0755)
	os.Mkdir("world/playerdata", 0755)
	os.Mkdir("world/poi", 0755)
	os.Mkdir("world/region", 0755)

	_, _ = os.Create("world/level.nbt")
	os.WriteFile("world/session.lock", nil, 0755)

	/*writer := gzip.NewWriter(lvl)
	f, _ := nbt.Marshal(data)
	writer.Write(f)
	lvl.Close()*/
}
