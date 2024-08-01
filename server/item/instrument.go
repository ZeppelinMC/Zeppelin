package item

type Instrument struct {
	SoundEvent  SoundEvent `nbt:"sound_event"`
	UseDuration int32      `nbt:"use_duration"`
	Range       float32    `nbt:"range"`
}

type SoundEvent struct {
	SoundID string  `nbt:"sound_id"`
	Range   float32 `nbt:"range"`
}
