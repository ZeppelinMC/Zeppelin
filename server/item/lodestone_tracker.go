package item

type LodestoneTracker struct {
	Target  Target `nbt:"target"`
	Tracked bool   `nbt:"tracked"`
}

type Target struct {
	Pos       []int32 `nbt:"pos"`
	Dimension string  `nbt:"dimension"`
}
