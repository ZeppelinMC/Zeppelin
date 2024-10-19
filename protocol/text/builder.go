package text

func New() TextComponent {
	return TextComponent{}
}

func (t TextComponent) WithColor(c string) TextComponent {
	t.Color = c
	return t
}

func (t TextComponent) WithText(c string) TextComponent {
	t.Text = c
	return t
}

func (t TextComponent) WithItalic() TextComponent {
	t.Italic = !t.Italic
	return t
}

func (t TextComponent) WithUnderline() TextComponent {
	t.Underlined = !t.Underlined
	return t
}

func (t TextComponent) WithStrikethrough() TextComponent {
	t.Strikethrough = !t.Strikethrough
	return t
}

func (t TextComponent) WithObfuscation() TextComponent {
	t.Obfuscated = !t.Obfuscated
	return t
}

func Color(c string) TextComponent {
	return TextComponent{Color: c}
}

func Text(c string) TextComponent {
	return TextComponent{Text: c}
}

func Italic() TextComponent {
	return TextComponent{Italic: true}
}

func Underline() TextComponent {
	return TextComponent{Underlined: true}
}

func Strikethrough() TextComponent {
	return TextComponent{Strikethrough: true}
}

func Obfuscation() TextComponent {
	return TextComponent{Obfuscated: true}
}
