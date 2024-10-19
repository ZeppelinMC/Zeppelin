// Package text provides encoding and decoding of text components
package text

import "fmt"

const (
	TypeText         = "text"
	TypeTranslatable = "translatable"
	TypeKeybind      = "keybind"
	TypeScore        = "score"
	TypeSelector     = "selector"
	TypeNBT          = "nbt"
)

const (
	Default    = "minecraft:default"
	Uniform    = "minecraft:uniform"
	Alt        = "minecraft:alt"
	Illageralt = "minecraft:illageralt"
)

const (
	OpenURL         = "open_url"
	RunCommand      = "run_command"
	SuggestCommand  = "suggest_command"
	ChangePage      = "change_page"
	CopyToClipboard = "copy_to_clipboard"
)

const (
	ShowText   = "show_text"
	ShowItem   = "show_item"
	ShowEntity = "show_entity"
)

type ClickEvent struct {
	Action string `json:"action" nbt:"action"`
	Value  string `json:"value" nbt:"value"`
}

type HoverEventContents struct {
	ID    string `json:"id,omitempty" nbt:"id,omitempty"`
	Count int    `json:"count,omitempty" nbt:"count,omitempty"`
	Tag   string `json:"tag,omitempty" nbt:"tag,omitempty"`
	Type  string `json:"type,omitempty" nbt:"type,omitempty"`
	Name  string `json:"name,omitempty" nbt:"name,omitempty"`
}

type HoverEvent struct {
	Action   string             `json:"action" nbt:"action"`
	Contents HoverEventContents `json:"contents" nbt:"contents"`
}

type TextComponent struct {
	Type  string          `json:"type,omitempty" nbt:"type,omitempty"`
	Extra []TextComponent `json:"extra,omitempty" nbt:"extra,omitempty"`

	Color         string     `json:"color,omitempty" nbt:"color,omitempty"`
	Bold          bool       `json:"bold,omitempty" nbt:"bold,omitempty"`
	Italic        bool       `json:"italic,omitempty" nbt:"italic,omitempty"`
	Underlined    bool       `json:"underlined,omitempty" nbt:"underlined,omitempty"`
	Strikethrough bool       `json:"strikethrough,omitempty" nbt:"strikethrough,omitempty"`
	Obfuscated    bool       `json:"obfuscated,omitempty" nbt:"obfuscated,omitempty"`
	Font          string     `json:"font,omitempty" nbt:"font,omitempty"`
	Insertion     string     `json:"insertion,omitempty" nbt:"insertion,omitempty"`
	ClickEvent    ClickEvent `json:"click_event,omitempty" nbt:"click_event,omitempty"`
	HoverEvent    HoverEvent `json:"hover_event,omitempty" nbt:"hover_event,omitempty"`

	Text string `json:"text" nbt:"text"`
}

func Sprint(a ...any) TextComponent {
	return TextComponent{Text: fmt.Sprint(a...)}
}

func Sprintf(format string, a ...any) TextComponent {
	return TextComponent{Text: fmt.Sprintf(format, a...)}
}

// Unmarshalf is the same as Unmarshal but with formatting
func Unmarshalf(codeChar rune, format string, v ...any) TextComponent {
	return Unmarshal(fmt.Sprintf(format, v...), codeChar)
}

// Unmarshal parses the text and color codes into a text component. The codeChar argument is what's used for color codes (i.e &)
func Unmarshal(text string, codeChar rune) TextComponent {
	var root TextComponent
	var component TextComponent
	var runeText = []rune(text)

	for i := 0; i < len(runeText); i++ {
		char := runeText[i]
		if char == codeChar && runeText[i+1] >= '0' && runeText[i+1] <= 'r' {
			if component.Text != "" {
				root.Extra = append(root.Extra, component)
				component = TextComponent{}
			}
			applyStyle(runeText[i+1], &component)
			i++
			continue
		}
		component.Text += string(char)
	}

	root.Extra = append(root.Extra, component)

	return root
}

// Marshal turns the text component to a classic color code message
func Marshal(component TextComponent, codeChar rune) string {
	var components = append([]TextComponent{component}, component.Extra...)

	var text string
	for _, comp := range components {
		componentStyles(comp, &text, codeChar)
		text += comp.Text
	}

	return text
}

func charColor(char rune) string {
	switch char {
	case '0':
		return "black"
	case '1':
		return "dark_blue"
	case '2':
		return "dark_green"
	case '3':
		return "dark_aqua"
	case '4':
		return "dark_red"
	case '5':
		return "dark_purple"
	case '6':
		return "gold"
	case '7':
		return "gray"
	case '8':
		return "dark_gray"
	case '9':
		return "blue"
	case 'a':
		return "green"
	case 'b':
		return "aqua"
	case 'c':
		return "red"
	case 'd':
		return "light_purple"
	case 'e':
		return "yellow"
	default: // case 'f'
		return "white"
	}
}

func colorChar(color string, char string) string {
	switch color {
	case "black":
		return char + "0"
	case "dark_blue":
		return char + "1"
	case "dark_green":
		return char + "2"
	case "dark_aqua":
		return char + "3"
	case "dark_red":
		return char + "4"
	case "dark_purple":
		return char + "5"
	case "gold":
		return char + "6"
	case "gray":
		return char + "7"
	case "dark_gray":
		return char + "8"
	case "blue":
		return char + "9"
	case "green":
		return char + "a"
	case "aqua":
		return char + "b"
	case "red":
		return char + "c"
	case "light_purple":
		return char + "d"
	case "yellow":
		return char + "e"
	default: // case "white"
		return char + "f"
	}
}

func applyStyle(char rune, component *TextComponent) {
	if char >= '0' && char <= 'f' {
		component.Color = charColor(char)
	}
	switch char {
	case 'k':
		component.Obfuscated = true
	case 'l':
		component.Bold = true
	case 'm':
		component.Strikethrough = true
	case 'n':
		component.Underlined = true
	case 'o':
		component.Italic = true
	}
}

func componentStyles(component TextComponent, text *string, codeChar rune) {
	*text += colorChar(component.Color, string(codeChar))
	if component.Obfuscated {
		*text += string(codeChar) + "k"
	}
	if component.Obfuscated {
		*text += string(codeChar) + "k"
	}
	if component.Bold {
		*text += string(codeChar) + "l"
	}
	if component.Strikethrough {
		*text += string(codeChar) + "m"
	}
	if component.Underlined {
		*text += string(codeChar) + "n"
	}
	if component.Italic {
		*text += string(codeChar) + "o"
	}
}
