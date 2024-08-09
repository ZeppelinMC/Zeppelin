package item

type WrittenBookContent struct {
	Pages      []Page `nbt:"pages"`
	Title      Title  `nbt:"title"`
	Author     string `nbt:"author"`
	Generation int32  `nbt:"generation"`
	Resolved   bool   `nbt:"resolved"`
}

type Page *struct { // ??????
	Raw              string `nbt:"raw"`
	Filtered         string `nbt:"filtered"`
	PlainTextContent string `nbt:"plain_text_content"` // ???
}

type Title struct {
	Raw      string `nbt:"raw"`
	Filtered string `nbt:"filtered"`
}
