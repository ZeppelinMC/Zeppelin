package world

import "strconv"

type DataPacks struct {
	Disabled []string `nbt:"Disabled"`
	Enabled  []string `nbt:"Enabled"`
}

type GameRule string

func (r GameRule) Int() (int, error) {
	return strconv.Atoi(string(r))
}

func (r GameRule) Bool() (bool, error) {
	return strconv.ParseBool(string(r))
}

/*type GameRules struct {
	DoDaylightCycle               GameRule `nbt:"doDaylightCycle"`
	DoInsomnia                    GameRule `nbt:"doInsomnia"`
	DoEntityDrops                 GameRule `nbt:"doEntityDrops"`
	SpectatorsGenerateChunks      GameRule `nbt:"spectatorsGenerateChunks"`
	DisableElytraMovementCheck    GameRule `nbt:"disableElytraMovementCheck"`
	DisableRaids                  GameRule `nbt:"disableRaids"`
	CommandModificationBlockLimit GameRule `nbt:"commandModificationBlockLimit"`
	ForgiveDeadPlayers            GameRule `nbt:"forgiveDeadPlayers"`
	MaxEntityCramming             GameRule `nbt:"maxEntityCramming"`
	SpawnRadius                   GameRule `nbt:"spawnRadius"`
	AnnounceAdvancements          GameRule `nbt:"announceAdvancements"`
	UniversalAnger                GameRule `nbt:"universalAnger"`
	FallDamage                    GameRule `nbt:"fallDamage"`
	RandomTickSpeed               GameRule `nbt:"randomTickSpeed"`
	DoTraderSpawning              GameRule `nbt:"doTraderSpawning"`
	DrowningDamage                GameRule `nbt:"drowningDamage"`
	ShowDeathMessages             GameRule `nbt:"showDeathMessages"`
	PlayersSleepingPercentage     GameRule `nbt:"playersSleepingPercentage"`
	DoImmediateRespawn            GameRule `nbt:"doImmediateRespawn"`
	NaturalRegeneration           GameRule `nbt:"naturalRegeneration"`
	KeepInventory                 GameRule `nbt:"keepInventory"`
	DoVinesSpread                 GameRule `nbt:"doVinesSpread"`
	DoTileDrops                   GameRule `nbt:"doTileDrops"`
	MaxCommandChainLength         GameRule `nbt:"maxCommandChainLength"`
	FireDamage                    GameRule `nbt:"fireDamage"`
	WaterSourceConversion         GameRule `nbt:"waterSourceConversion"`
	CommandBlockOutput            GameRule `nbt:"commandBlockOutput"`
	DoWeatherCycle                GameRule `nbt:"doWeatherCycle"`
	SnowAccumulationHeight        GameRule `nbt:"snowAccumulationHeight"`
	DoWardenSpawning              GameRule `nbt:"doWardenSpawning"`
	DoFireTick                    GameRule `nbt:"doFireTick"`
	DoPatrolSpawning              GameRule `nbt:"doPatrolSpawning"`
	DoMobLoot                     GameRule `nbt:"doMobLoot"`
	SendCommandFeedback           GameRule `nbt:"sendCommandFeedback"`
	DoMobSpawning                 GameRule `nbt:"doMobSpawning"`
	MobExplosionDropDecay         GameRule `nbt:"mobExplosionDropDecay"`
	MobGriefing                   GameRule `nbt:"mobGriefing"`
	FreezeDamage                  GameRule `nbt:"freezeDamage"`
	LogAdminCommands              GameRule `nbt:"logAdminCommands"`
	GlobalSoundEvents             GameRule `nbt:"globalSoundEvents"`
	TntExplosionDropDecay         GameRule `nbt:"tntExplosionDropDecay"`
	BlockExplosionDropDecay       GameRule `nbt:"blockExplosionDropDecay"`
	DoLimitedCrafting             GameRule `nbt:"doLimitedCrafting"`
	ReducedDebugInfo              GameRule `nbt:"reducedDebugInfo"`
	LavaSourceConversion          GameRule `nbt:"lavaSourceConversion"`
}*/

type DragonFight struct {
	NeedsStateScanning int8    `nbt:"NeedsStateScanning"`
	Gateways           []int32 `nbt:"Gateways"`
	DragonKilled       int8    `nbt:"DragonKilled"`
	PreviouslyKilled   int8    `nbt:"PreviouslyKilled"`
}

type Version struct {
	Snapshot int8   `nbt:"Snapshot"`
	Series   string `nbt:"Series"`
	Id       int32  `nbt:"Id"`
	Name     string `nbt:"Name"`
}

type BiomeSource struct {
	Preset string `nbt:"preset"`
	Type   string `nbt:"type"`
}

type DimensionGeneratorLayer struct {
	Block  string `nbt:"block"`
	Height int32  `nbt:"height"`
}

type DimensionGeneratorSettings struct {
	Biome              string                    `nbt:"biome"`
	Features           int32                     `nbt:"features"`
	Lakes              int32                     `nbt:"lakes"`
	Layers             []DimensionGeneratorLayer `nbt:"layers"`
	StructureOverrides []string                  `nbt:"structure_overrides"`
}

type DimensionGenerator struct {
	BiomeSource BiomeSource `nbt:"biome_source"`
	Type        string      `nbt:"type"`
	//Settings    string `nbt:"settings"` // string or DimensionGeneratorSettings
}

type DimensionData struct {
	Generator DimensionGenerator `nbt:"generator"`
	Type      string             `nbt:"type"`
}

type Dimensions struct {
	End       DimensionData `nbt:"minecraft:the_end"`
	Overworld DimensionData `nbt:"minecraft:overworld"`
	Nether    DimensionData `nbt:"minecraft:the_nether"`
}

type WorldGenSettings struct {
	BonusChest       int8       `nbt:"bonus_chest"`
	Seed             int64      `nbt:"seed"`
	GenerateFeatures int8       `nbt:"generate_features"`
	Dimensions       Dimensions `nbt:"dimensions"`
}

type WorldData struct {
	ServerBrands               []string            `nbt:"ServerBrands"`
	SpawnY                     int32               `nbt:"SpawnY"`
	ThunderTime                int32               `nbt:"thunderTime"`
	LastPlayed                 int64               `nbt:"LastPlayed"`
	BorderSize                 float64             `nbt:"BorderSize"`
	DataVersion                int32               `nbt:"DataVersion"`
	Time                       int64               `nbt:"Time"`
	Difficulty                 int8                `nbt:"Difficulty"`
	Raining                    int8                `nbt:"raining"`
	BorderWarningBlocks        float64             `nbt:"BorderWarningBlocks"`
	WorldGenSettings           WorldGenSettings    `nbt:"WorldGenSettings"`
	BorderSizeLerpTarget       float64             `nbt:"BorderSizeLerpTarget"`
	Version                    Version             `nbt:"Version"`
	DayTime                    int64               `nbt:"DayTime"`
	WanderingTraderSpawnChance int32               `nbt:"WanderingTraderSpawnChance"`
	Hardcore                   int8                `nbt:"hardcore"`
	SpawnX                     int32               `nbt:"SpawnX"`
	BorderWarningTime          float64             `nbt:"BorderWarningTime"`
	WanderingTraderSpawnDelay  int32               `nbt:"WanderingTraderSpawnDelay"`
	CustomBossEvents           struct{}            `nbt:"CustomBossEvents"`
	Player                     PlayerData          `nbt:"Player"`
	ClearWeatherTime           int32               `nbt:"clearWeatherTime"`
	Thundering                 int8                `nbt:"thundering"`
	SpawnAngle                 float32             `nbt:"SpawnAngle"`
	ScheduledEvents            []interface{}       `nbt:"ScheduledEvents"`
	BorderCenterZ              float64             `nbt:"BorderCenterZ"`
	RainTime                   int32               `nbt:"rainTime"`
	LevelName                  string              `nbt:"LevelName"`
	WasModded                  int8                `nbt:"WasModded"`
	BorderCenterX              float64             `nbt:"BorderCenterX"`
	AllowCommands              int8                `nbt:"allowCommands"`
	DragonFight                DragonFight         `nbt:"DragonFight"`
	BorderDamagePerBlock       float64             `nbt:"BorderDamagePerBlock"`
	Initialized                int8                `nbt:"initialized"`
	GameRules                  map[string]GameRule `nbt:"GameRules"`
	SpawnZ                     int32               `nbt:"SpawnZ"`
	BorderSafeZone             float64             `nbt:"BorderSafeZone"`
	DataPacks                  DataPacks           `nbt:"DataPacks"`
	GameType                   int32               `nbt:"GameType"`
	DifficultyLocked           int8                `nbt:"DifficultyLocked"`
	BorderSizeLerpTime         int64               `nbt:"BorderSizeLerpTime"`
	VersionNumber              int32               `nbt:"version"`
}

type worldData struct {
	Data WorldData `nbt:"Data"`
}
