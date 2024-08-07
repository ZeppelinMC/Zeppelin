package item

type DyedColor struct {
	RGB           int32 `nbt:"rgb"`
	ShowInTooltip bool  `nbt:"show_in_tooltip"`
}
