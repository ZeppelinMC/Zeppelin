package item

type Trim struct {
	Pattern       string `nbt:"pattern"`
	Material      string `nbt:"material"`
	ShowInTooltip bool   `nbt:"show_in_tooltip"`
}
