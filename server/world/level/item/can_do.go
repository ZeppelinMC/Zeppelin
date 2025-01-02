package item

// Both CanBreak and CanPlace have the same contents so im using the same class
type CanDo struct {
	ShowInTooltip bool         `nbt:"show_in_tooltip"`
	State         any          `nbt:"state"`  // Have to put any cause theres an infinite amount of key value pairs there could be
	Blocks        any          `nbt:"blocks"` //Can be string, or list
	NBT           string       `nbt:"nbt"`
	Predicates    []Predicates `nbt:"predicates"`
}

type Predicates struct {
	State  any    `nbt:"state"`
	Blocks any    `nbt:"blocks"`
	NBT    string `nbt:"nbt"`
}
