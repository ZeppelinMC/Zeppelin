package item

type CanBreak struct {
	ShowInTooltip bool   `nbt:"show_in_tooltip"`
	Blocks        *int32 `nbt:"blocks"`
	NBT           *int32 `nbt:"nbt"`
}

type Predicate []struct {
	Blocks *string `nbt:"blocks"`
	NBT    int32   `nbt:"nbt"`
	State  int32   `nbt:"state"`
}

type State struct {
	Name  string `nbt:"name"`
	Value string `nbt:"value"` // ????????
}
