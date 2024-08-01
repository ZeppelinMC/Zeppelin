package item

type MapDecorations struct { // I'm a little confused on this one. It goes <string> for the head tag??
	Type     string  `nbt:"type"`
	X        float32 `nbt:"x"`
	Z        float32 `nbt:"z"`
	Rotation float32 `nbt:"rotation"`
}
