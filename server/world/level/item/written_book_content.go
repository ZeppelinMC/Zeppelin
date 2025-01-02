package item

type WrittenBookContent struct {
	Pages      any    `nbt:"pages"`
	Title      Title  `nbt:"title"`
	Author     string `nbt:"author"`
	Generation int32  `nbt:"generation"`
	Resolved   bool   `nbt:"resolved"`
}

type Title struct {
	Raw      string `nbt:"raw"`
	Filtered string `nbt:"filtered"`
}
