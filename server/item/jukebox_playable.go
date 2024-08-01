package item

type JukeboxPlayable struct {
	Song          string `nbt:"song"`
	ShowInTooltip bool   `nbt:"show_in_tooltip"`
}
