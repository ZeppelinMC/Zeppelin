package item

type Food struct {
	Nutrition       int32    `nbt:"nutrition"`
	Saturation      float32  `nbt:"saturation"`
	CanAlwaysEat    bool     `nbt:"can_always_eat"`
	EatSeconds      float32  `nbt:"eat_seconds"`
	UsingConvertsTo Item     `nbt:"using_converts_to"`
	Effects         []Effect `nbt:"effects"`
}

type Effects struct {
	Effect      Effect  `nbt:"effect"`
	Probability float32 `nbt:"probability"`
}

type Effect struct {
	ID            string `nbt:"id"`
	Amplifier     int8   `nbt:"amplifier"`
	Duration      int32  `nbt:"duration"`
	Ambient       bool   `nbt:"ambient"`
	ShowParticles bool   `nbt:"show_particles"`
	ShowIcon      bool   `nbt:"show_icon"`
}
