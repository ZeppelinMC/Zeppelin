package qnbt

import (
	_ "embed"
	"os"
	"runtime/pprof"
	"testing"
)

//go:embed testdata/chunk.nbt
var chunkData []byte

type chunk struct {
	Status        string
	ZPos          int32 `nbt:"zPos"`
	LastUpdate    int64
	InhabitedTime int64
	XPos          int32 `nbt:"xPos"`
	YPos          int32 `nbt:"yPos"`
	Heightmaps    struct {
		OceanFloor             []int64 `nbt:"OCEAN_FLOOR"`
		MotionBlockingNoLeaves []int64 `nbt:"MOTION_BLOCKING_NO_LEAVES"`
		MotionBlocking         []int64 `nbt:"MOTION_BLOCKING"`
		WorldSurface           []int64 `nbt:"WORLD_SURFACE"`
	}
	BlockEntities []struct {
		Id         string `nbt:"id"`
		KeepPacked bool   `nbt:"keepPacked"`
		Normal     struct {
			SimultaneousMobs               float32 `nbt:"simultaneous_mobs"`
			SimultaneousMobsAddedPerPlayer float32 `nbt:"simultaneous_mobs_added_per_player"`
			SpawnPotentials                []struct {
				Data struct {
					Entity struct {
						Id string `nbt:"id"`
					} `nbt:"entity"`
				} `nbt:"data"`
				Weight int32 `nbt:"weight"`
			} `nbt:"spawn_potentials"`
			TicksBetweenSpawn int32 `nbt:"ticks_between_spawn"`
		} `nbt:"normal_config"`
		Ominous struct {
			Loot []struct {
				Data   string `nbt:"data"`
				Weight int32  `nbt:"weight"`
			} `nbt:"loot_tables_to_eject"`

			SimultaneousMobs               float32 `nbt:"simultaneous_mobs"`
			SimultaneousMobsAddedPerPlayer float32 `nbt:"simultaneous_mobs_added_per_player"`
			TicksBetweenSpawn              int32   `nbt:"ticks_between_spawn"`

			SpawnPotentials []struct {
				Data struct {
					Entity struct {
						Id string `nbt:"id"`
					} `nbt:"entity"`
					Equipment struct {
						LootTable       string  `nbt:"loot_table"`
						SlotDropChances float32 `nbt:"slot_drop_chances"`
					} `nbt:"equipment"`
				} `nbt:"data"`
				Weight int32 `nbt:"weight"`
			} `nbt:"spawn_potentials"`
		} `nbt:"ominous_config"`
		X int32 `nbt:"x"`
		Y int32 `nbt:"y"`
		Z int32 `nbt:"z"`
	} `nbt:"block_entities"`
	BlockTicks []struct {
		I string `nbt:"i"`
		P int32  `nbt:"p"`
		T int32  `nbt:"t"`
		X int32  `nbt:"x"`
		Y int32  `nbt:"y"`
		Z int32  `nbt:"z"`
	} `nbt:"block_ticks"`
	PostProcessing [][]int16
	IsLightOn      int8 `nbt:"isLightOn"`
	Sections       []struct {
		Y                    int8
		BlockLight, SkyLight []byte
		Biomes               struct {
			Data    []int64  `nbt:"data"`
			Palette []string `nbt:"palette"`
		} `nbt:"biomes"`
		BlockStates struct {
			Data    []int64 `nbt:"data"`
			Palette []struct {
				Name       string
				Properties map[string]string
			} `nbt:"palette"`
		} `nbt:"block_states"`
	} `nbt:"sections"`

	Structures struct {
		References struct {
			Tc []int64 `nbt:"minecraft:trial_chambers"`
		}
		Starts struct{} `nbt:"starts"`
	} `nbt:"structures"`

	DataVersion int32
}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshal(chunkData, &chunk{})
	}
	b.ReportAllocs()
	f, _ := os.Create("testdata/mem.prof")
	pprof.Lookup("allocs").WriteTo(f, 0)
	f.Close()
}
