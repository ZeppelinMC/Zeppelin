package metadata

// Player returns the default player metadata object
func Player(health float32) Metadata {
	return Metadata{
		// Entity
		BaseIndex:                      Byte(0),
		AirTicksIndex:                  VarInt(300),
		CustomNameIndex:                OptionalTextComponent(nil),
		IsCustomNameVisibleIndex:       Boolean(false),
		IsSilentIndex:                  Boolean(false),
		HasNoGravityIndex:              Boolean(false),
		PoseIndex:                      Standing,
		TicksFrozenInPowderedSnowIndex: VarInt(0),
		// Living Entity extends Entity
		LivingEntityHandstatesIndex: Byte(0),
		LivingEntityHealthIndex:     Float(health), // this one is actually stored in data (others are too but arent used yet)
		//LivingEntityPotionEffectColorIndex:   VarInt(0),
		LivingEntityPotionEffectAmbientIndex: Boolean(false),
		LivingEntityArrowCountIndex:          VarInt(0),
		LivingEntityBeeStingersCountIndex:    VarInt(0),
		LivingEntitySleepingBedPositionIndex: OptionalPosition(nil),
		// Player extends Living Entity
		PlayerAdditionalHeartsIndex:   Float(0),
		PlayerScoreIndex:              VarInt(0),
		PlayerDisplayedSkinPartsIndex: Byte(0),
		PlayerMainHandIndex:           Byte(1),
	}
}
