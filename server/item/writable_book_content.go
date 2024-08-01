package item

type WritableBookContent struct {
	Pages []BookPage `nbt:"pages"`
}

type BookPage *struct { // ??????
	Raw              string `nbt:"raw"`
	Filtered         string `nbt:"filtered"`
	PlainTextContent string `nbt:"plain_text_content"` // ???
}
