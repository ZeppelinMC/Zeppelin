package item

type Enchantments struct {
	ShowInTooltip bool `nbt:"show_in_tooltip"`
	Levels        any  `nbt:"levels"`
}
