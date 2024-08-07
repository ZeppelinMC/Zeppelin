package item

type Container *struct {
	Slot int32 `nbt:"slot"`
}

type item struct {
	ID    string `nbt:"id"`
	Count int32  `nbt:"count"`
	// components TAG????
}
