package item

type PotionContents struct {
	Portion       string          `nbt:"potion"`
	CustomColor   int32           `nbt:"custom_color"`
	CustomEffects []CustomEffects `nbt:"custom_effects"`
}

type CustomEffects *struct {
	ID            string `nbt:"id"`
	Amplifier     byte   `nbt:"byte"`
	Duration      int32  `nbt:"duration"`
	Ambient       bool   `nbt:"ambient"`
	ShowParticles bool   `nbt:"show_particles"`
	ShowIcon      bool   `nbt:"show_icon"`
}
