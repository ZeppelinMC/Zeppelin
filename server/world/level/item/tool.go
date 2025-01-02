package item

type Tool struct {
	DefaultMiningSpeed float32 `nbt:"default_mining_speed"`
	DamagePerBlock     int32   `nbt:"damage_per_block"`
	Rules              []Rule  `nbt:"rules"`
}

type Rule struct {
	Blocks          any     `nbt:"blocks"`
	Speed           float32 `nbt:"speed"`
	CorrectForDrops bool    `nbt:"correct_for_drops"`
}
