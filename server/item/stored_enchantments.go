package item

type StoredEnchantments struct {
	Levels        Level `nbt:"levels"`
	ShowInTooltip bool  `nbt:"show_in_tooltip"`
}

type Level struct {
	EnchantmentID int32 `nbt:"enchantment_id"` // ???
}
