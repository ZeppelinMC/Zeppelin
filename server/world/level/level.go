// Package level provides documentation and loading of .dat files found in the world folder (according to https://minecraft.wiki)
// .dat files are NBT (Named Binary Tag) files compressed with Gunzip.
// when writing a .dat file, the previous .dat file is renamed to .dat_old

package level

import (
	"compress/gzip"
	"errors"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/zeppelinmc/zeppelin/protocol/nbt"
	"github.com/zeppelinmc/zeppelin/protocol/properties"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/level/seed"
)

var ErrAlreadyClosed = errors.New("already closed")

var SessionLock = []byte{0xE2, 0x98, 0x83}

// A game rule is a string containing an integer or a boolean
type GameRule string

// Returns the boolean inside of the game rule string
func (rule GameRule) Boolean() (bool, error) {
	return strconv.ParseBool(string(rule))
}

// Returns the int inside of the game rule string
func (rule GameRule) Integer() (int, error) {
	return strconv.Atoi(string(rule))
}

// A date represented by the amount of milliseconds since January 1st, 1970
type UnixMilliTimestamp int64

func (stamp UnixMilliTimestamp) Time() time.Time {
	return time.UnixMilli(int64(stamp))
}

func Now() UnixMilliTimestamp {
	return UnixMilliTimestamp(time.Now().UnixMilli())
}

type DimensionGenerationSettings struct {
	Generator struct {
		BiomeSource struct {
			Preset string `nbt:"preset,omitempty"`
			Type   string `nbt:"type"`
		} `nbt:"biome_source"`
		//Settings string `nbt:"settings"` // Can be both string and compound, skip for now
		Type string `nbt:"type"`
	} `nbt:"generator"`
	Type string `nbt:"type"`
}

type Level struct {
	Data struct {
		BorderCenterX        float64
		BorderCenterZ        float64
		BorderDamagePerBlock float64
		BorderSize           float64
		BorderSafeZone       float64
		BorderSizeLerpTarget float64
		BorderSizeLerpTime   int64
		BorderWarningBlocks  float64
		BorderWarningTime    float64

		DataPacks struct {
			Disabled []string
			Enabled  []string
		}

		DataVersion      int32
		DayTime          int64
		Difficulty       byte
		DifficultyLocked bool
		DragonFight      struct {
			DragonKilled       bool
			Gateways           []int32
			NeedsStateScanning bool
			PreviouslyKilled   byte
		}

		GameRules    map[string]GameRule
		GameType     GameMode
		LastPlayed   UnixMilliTimestamp
		LevelName    string
		ServerBrands []string

		SpawnAngle             float32
		SpawnX, SpawnY, SpawnZ int32

		Time    int64 // time since the world has started in ticks
		Version struct {
			Id       int32
			Name     string
			Series   string
			Snapshot int8
		}

		WanderingTraderId          []int32
		WanderingTraderSpawnChance int32
		WanderingTraderSpawnDelay  int32
		WasModded                  bool

		WorldGenSettings struct {
			BonusChest       bool                                   `nbt:"bonus_chest"`
			Dimensions       map[string]DimensionGenerationSettings `nbt:"dimensions"`
			GenerateFeatures bool                                   `nbt:"generate_features"`
			Seed             seed.Seed                              `nbt:"seed"`
		}

		AllowCommands    bool  `nbt:"allowCommands"`
		ClearWeatherTime int32 `nbt:"clearWeatherTime"`
		Hardcore         bool  `nbt:"hardcore"`
		Initialized      bool  `nbt:"initialized"`

		RainTime    int32 `nbt:"rainTime"`
		Raining     bool  `nbt:"raining"`
		Thundertime int32 `nbt:"thunderTime"`
		Thundering  bool  `nbt:"thundering"`
		VersionInt  int32 `nbt:"version"`
	}

	// the base path of the world
	basePath string `nbt:"-"`

	closed bool
}

// worldPath is the base path of the world
func Open(worldPath string) (Level, error) {
	var level Level
	file, err := os.Open(worldPath + "/level.dat")
	if err != nil {
		return level, err
	}
	rd, err := gzip.NewReader(file)
	if err != nil {
		return level, err
	}

	var buf, _ = io.ReadAll(rd)

	rd.Close()
	file.Close()

	_, err = nbt.Unmarshal(buf, &level)
	level.basePath = worldPath

	return level, err
}

func Create(l Level) error {
	file, err := os.Create(l.basePath + "/level.dat")
	if err != nil {
		return err
	}
	w := gzip.NewWriter(file)

	defer w.Close()
	defer file.Close()

	return nbt.NewEncoder(w).Encode("", l)
}

func (l *Level) Close() error {
	if l.closed {
		return ErrAlreadyClosed
	}
	l.closed = true
	return Create(*l)
}

func New(gen chunk.Generator, props properties.ServerProperties, worldPath string) Level {
	var l Level
	l.Data.SpawnX, l.Data.SpawnY, l.Data.SpawnZ = gen.GenerateWorldSpawn()
	l.Data.AllowCommands = true
	l.Data.BorderDamagePerBlock = 0.2
	l.Data.BorderSize = 60000000
	l.Data.BorderSafeZone = 5
	l.Data.BorderSizeLerpTarget = 60000000
	l.Data.BorderWarningBlocks = 5
	l.Data.BorderWarningTime = 15
	l.Data.DataPacks.Enabled = []string{"vanilla"}
	l.Data.Difficulty = diffstr(props.Difficulty)
	l.Data.WorldGenSettings.Seed = seed.New(props.LevelSeed)
	l.Data.WorldGenSettings.GenerateFeatures = props.GenerateStructures
	l.Data.GameType = gmstr(props.Gamemode)
	l.Data.Hardcore = props.Hardcore
	l.Data.Initialized = true
	l.Data.LastPlayed = Now()
	l.Data.LevelName = props.LevelName
	l.Data.VersionInt = 19133

	l.Data.Version.Id = 3953
	l.Data.Version.Name = "1.21"
	l.Data.Version.Series = "main"
	l.Data.ServerBrands = []string{"Zeppelin"}
	l.basePath = worldPath

	return l
}

func (l *Level) Refresh(props properties.ServerProperties) {
	l.Data.AllowCommands = true
	l.Data.DataPacks.Enabled = []string{"vanilla"}
	l.Data.Difficulty = diffstr(props.Difficulty)
	l.Data.WorldGenSettings.GenerateFeatures = props.GenerateStructures
	l.Data.GameType = gmstr(props.Gamemode)
	l.Data.Hardcore = props.Hardcore
	l.Data.Initialized = true
	l.Data.LastPlayed = Now()
	l.Data.LevelName = props.LevelName
	l.Data.VersionInt = 19133

	l.Data.Version.Id = 3953
	l.Data.Version.Name = "1.21"
	l.Data.Version.Series = "main"
	l.Data.ServerBrands = []string{"Zeppelin"}
}

func diffstr(str string) byte {
	switch str {
	case "peaceful":
		return 0
	case "normal":
		return 2
	case "hard":
		return 3
	default:
		return 1
	}
}

func gmstr(str string) GameMode {
	switch str {
	case "creative":
		return GameModeCreative
	case "adventure":
		return GameModeAdventure
	case "spectator":
		return GameModeSpectator

	default:
		return GameModeSurvival
	}
}
