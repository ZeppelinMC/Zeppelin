package item

type Enchantments struct {
	ShowInTooltip bool `nbt:"show_in_tooltip"`
}

type Levels struct {
	EnchantmentID int32 `nbt:"enchantment_id"` // ????????
}
