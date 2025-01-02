package item

type Fireworks struct {
	FlightDuration int8 `nbt:"flight_duration"`
	Explosions []Explosion `nbt:"explosions"`
}

type Explosion struct {
	Shape      string  `nbt:"shape"`
	Colors     []int32 `nbt:"colors"`
	FadeColors []int32 `nbt:"fade_colors"`
	HasTrail   bool    `nbt:"has_trail"`
	HasTwinkle bool    `nbt:"has_twinkle"`
}
