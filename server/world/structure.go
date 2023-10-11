package world

import "strconv"

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
