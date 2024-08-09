package item

type Food struct {
	Nutrition    int32   `nbt:"nutrition"`
	Saturation   float32 `nbt:"saturation"`
	CanAlwaysEat bool    `nbt:"can_always_eat"`
	EatSeconds   float32 `nbt:"eat_seconds"`
}

type UsingConvertsTo struct {
	ID string `nbt:"id"`
	// components TAG ????
}

type Effect *struct {
	ID            string  `nbt:"id"`
	Amplifier     byte    `nbt:"amplifier"`
	Duration      int32   `nbt:"duration"`
	Ambient       bool    `nbt:"ambient"`
	ShowParticles bool    `nbt:"show_particles"`
	ShowIcon      bool    `nbt:"show_icon"`
	Probability   float32 `nbt:"probability"`
}
