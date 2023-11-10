package color

import (
	"fmt"
	"strings"

	"github.com/aimjel/minecraft/chat"
)

type Attribute = string
type Color []Attribute

const (
	Reset         Attribute = "0"
	Bold          Attribute = "1"
	Dim           Attribute = "2"
	Italic        Attribute = "3"
	Underline     Attribute = "4"
	Blinking      Attribute = "5"
	Inverse       Attribute = "6"
	Hidden        Attribute = "7"
	Strikethrough Attribute = "8"

	FgBlack   Attribute = "30"
	FgRed     Attribute = "31"
	FgGreen   Attribute = "32"
	FgYellow  Attribute = "33"
	FgBlue    Attribute = "34"
	FgMagenta Attribute = "35"
	FgCyan    Attribute = "36"
	FgWhite   Attribute = "37"

	BgBlack   Attribute = "40"
	BgRed     Attribute = "41"
	BgGreen   Attribute = "42"
	BgYellow  Attribute = "43"
	BgBlue    Attribute = "44"
	BgMagenta Attribute = "45"
	BgCyan    Attribute = "46"
	BgWhite   Attribute = "47"

	FgBrBlack   Attribute = "90"
	FgBrRed     Attribute = "91"
	FgBrGreen   Attribute = "92"
	FgBrYellow  Attribute = "93"
	FgBrBlue    Attribute = "94"
	FgBrMagenta Attribute = "95"
	FgBrCyan    Attribute = "96"
	FgBrWhite   Attribute = "97"

	BgBrBlack   Attribute = "100"
	BgBrRed     Attribute = "101"
	BgBrGreen   Attribute = "102"
	BgBrYellow  Attribute = "103"
	BgBrBlue    Attribute = "104"
	BgBrMagenta Attribute = "105"
	BgBrCyan    Attribute = "106"
	BgBrWhite   Attribute = "107"
)

const escape = "\x1b"

func (c Color) Colorize(str string) string {
	return fmt.Sprintf("%s[%sm%s%s[0m", escape, strings.Join(c, ";"), str, escape)
}

func FromChat(msg chat.Message) string {
	var str strings.Builder
	texts := []chat.Message{msg}
	for _, m := range msg.Extra {
		texts = append(texts, m)
		texts = append(texts, m.Extra...)
	}

	for _, text := range texts {
		if text.Text == nil {
			continue
		}
		var color = Color{chatColors[text.Color]}
		if text.Bold {
			color = append(color, Bold)
		}
		if text.Italic {
			color = append(color, Italic)
		}
		if text.Underlined {
			color = append(color, Underline)
		}
		if text.Strikethrough {
			color = append(color, Strikethrough)
		}
		str.WriteString(color.Colorize(*text.Text))
	}

	return strings.ReplaceAll(str.String(), "\n", "\n\r")
}

var chatColors = map[string]Attribute{
	"black":        FgBlack,
	"dark_blue":    FgBlue,
	"dark_green":   FgGreen,
	"dark_aqua":    FgCyan,
	"dark_red":     FgRed,
	"dark_purple":  FgMagenta,
	"gold":         FgYellow,
	"gray":         FgWhite,
	"dark_gray":    FgBlack,
	"blue":         FgBlue,
	"green":        FgGreen,
	"aqua":         FgCyan,
	"red":          FgRed,
	"light_purple": FgMagenta,
	"yellow":       FgYellow,
	"white":        FgWhite,
}
