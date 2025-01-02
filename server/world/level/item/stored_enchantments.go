package item

type StoredEnchantments struct {
	Levels        any  `nbt:"levels"`
	ShowInTooltip bool `nbt:"show_in_tooltip"`
}
