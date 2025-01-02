package item

type FireworkExplosion struct {
	Shape      string  `nbt:"shape"`
	Colors     []int32 `nbt:"colors"`
	FadeColors []int32 `nbt:"fade_colors"`
	HasTrail   bool    `nbt:"has_trail"`
	HasTwinkle bool    `nbt:"has_twinkle"`
}
