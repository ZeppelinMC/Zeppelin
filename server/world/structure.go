package world

type DataPacks struct {
	Disabled []string `nbt:"Disabled"`
	Enabled  []string `nbt:"Enabled"`
}

type GameRules struct {
	DoDaylightCycle               string `nbt:"doDaylightCycle"`
	DoInsomnia                    string `nbt:"doInsomnia"`
	DoEntityDrops                 string `nbt:"doEntityDrops"`
	SpectatorsGenerateChunks      string `nbt:"spectatorsGenerateChunks"`
	DisableElytraMovementCheck    string `nbt:"disableElytraMovementCheck"`
	DisableRaids                  string `nbt:"disableRaids"`
	CommandModificationBlockLimit string `nbt:"commandModificationBlockLimit"`
	ForgiveDeadPlayers            string `nbt:"forgiveDeadPlayers"`
	MaxEntityCramming             string `nbt:"maxEntityCramming"`
	SpawnRadius                   string `nbt:"spawnRadius"`
	AnnounceAdvancements          string `nbt:"announceAdvancements"`
	UniversalAnger                string `nbt:"universalAnger"`
	FallDamage                    string `nbt:"fallDamage"`
	RandomTickSpeed               string `nbt:"randomTickSpeed"`
	DoTraderSpawning              string `nbt:"doTraderSpawning"`
	DrowningDamage                string `nbt:"drowningDamage"`
	ShowDeathMessages             string `nbt:"showDeathMessages"`
	PlayersSleepingPercentage     string `nbt:"playersSleepingPercentage"`
	DoImmediateRespawn            string `nbt:"doImmediateRespawn"`
	NaturalRegeneration           string `nbt:"naturalRegeneration"`
	KeepInventory                 string `nbt:"keepInventory"`
	DoVinesSpread                 string `nbt:"doVinesSpread"`
	DoTileDrops                   string `nbt:"doTileDrops"`
	MaxCommandChainLength         string `nbt:"maxCommandChainLength"`
	FireDamage                    string `nbt:"fireDamage"`
	WaterSourceConversion         string `nbt:"waterSourceConversion"`
	CommandBlockOutput            string `nbt:"commandBlockOutput"`
	DoWeatherCycle                string `nbt:"doWeatherCycle"`
	SnowAccumulationHeight        string `nbt:"snowAccumulationHeight"`
	DoWardenSpawning              string `nbt:"doWardenSpawning"`
	DoFireTick                    string `nbt:"doFireTick"`
	DoPatrolSpawning              string `nbt:"doPatrolSpawning"`
	DoMobLoot                     string `nbt:"doMobLoot"`
	SendCommandFeedback           string `nbt:"sendCommandFeedback"`
	DoMobSpawning                 string `nbt:"doMobSpawning"`
	MobExplosionDropDecay         string `nbt:"mobExplosionDropDecay"`
	MobGriefing                   string `nbt:"mobGriefing"`
	FreezeDamage                  string `nbt:"freezeDamage"`
	LogAdminCommands              string `nbt:"logAdminCommands"`
	GlobalSoundEvents             string `nbt:"globalSoundEvents"`
	TntExplosionDropDecay         string `nbt:"tntExplosionDropDecay"`
	BlockExplosionDropDecay       string `nbt:"blockExplosionDropDecay"`
	DoLimitedCrafting             string `nbt:"doLimitedCrafting"`
	ReducedDebugInfo              string `nbt:"reducedDebugInfo"`
	LavaSourceConversion          string `nbt:"lavaSourceConversion"`
}

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

type DimensionGenerator struct {
	BiomeSource BiomeSource `nbt:"biome_source"`
	Type        string      `nbt:"type"`
	Settings    string      `nbt:"settings"`
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
	ServerBrands               []string         `nbt:"ServerBrands"`
	SpawnY                     int32            `nbt:"SpawnY"`
	ThunderTime                int32            `nbt:"thunderTime"`
	LastPlayed                 int64            `nbt:"LastPlayed"`
	BorderSize                 float64          `nbt:"BorderSize"`
	DataVersion                int32            `nbt:"DataVersion"`
	Time                       int64            `nbt:"Time"`
	Difficulty                 int8             `nbt:"Difficulty"`
	Raining                    int8             `nbt:"raining"`
	BorderWarningBlocks        float64          `nbt:"BorderWarningBlocks"`
	WorldGenSettings           WorldGenSettings `nbt:"WorldGenSettings"`
	BorderSizeLerpTarget       float64          `nbt:"BorderSizeLerpTarget"`
	Version                    Version          `nbt:"Version"`
	DayTime                    int64            `nbt:"DayTime"`
	WanderingTraderSpawnChance int32            `nbt:"WanderingTraderSpawnChance"`
	Hardcore                   int8             `nbt:"hardcore"`
	SpawnX                     int32            `nbt:"SpawnX"`
	BorderWarningTime          float64          `nbt:"BorderWarningTime"`
	WanderingTraderSpawnDelay  int32            `nbt:"WanderingTraderSpawnDelay"`
	CustomBossEvents           struct{}         `nbt:"CustomBossEvents"`
	Player                     PlayerData       `nbt:"Player"`
	ClearWeatherTime           int32            `nbt:"clearWeatherTime"`
	Thundering                 int8             `nbt:"thundering"`
	SpawnAngle                 float32          `nbt:"SpawnAngle"`
	ScheduledEvents            []interface{}    `nbt:"ScheduledEvents"`
	BorderCenterZ              float64          `nbt:"BorderCenterZ"`
	RainTime                   int32            `nbt:"rainTime"`
	LevelName                  string           `nbt:"LevelName"`
	WasModded                  int8             `nbt:"WasModded"`
	BorderCenterX              float64          `nbt:"BorderCenterX"`
	AllowCommands              int8             `nbt:"allowCommands"`
	DragonFight                DragonFight      `nbt:"DragonFight"`
	BorderDamagePerBlock       float64          `nbt:"BorderDamagePerBlock"`
	Initialized                int8             `nbt:"initialized"`
	GameRules                  GameRules        `nbt:"GameRules"`
	SpawnZ                     int32            `nbt:"SpawnZ"`
	BorderSafeZone             float64          `nbt:"BorderSafeZone"`
	DataPacks                  DataPacks        `nbt:"DataPacks"`
	GameType                   int32            `nbt:"GameType"`
	DifficultyLocked           int8             `nbt:"DifficultyLocked"`
	BorderSizeLerpTime         int64            `nbt:"BorderSizeLerpTime"`
	VersionNumber              int32            `nbt:"version"`
}

type worldData struct {
	Data WorldData `nbt:"Data"`
}
