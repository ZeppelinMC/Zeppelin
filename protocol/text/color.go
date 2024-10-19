package text

import (
	"fmt"
	"image/color"
)

// colors
const (
	Black       = "black"
	DarkBlue    = "dark_blue"
	DarkGreen   = "dark_green"
	DarkCyan    = "dark_aqua"
	DarkRed     = "dark_red"
	Purple      = "dark_purple"
	Gold        = "gold"
	Gray        = "gray"
	DarkGray    = "dark_gray"
	Blue        = "blue"
	BrightGreen = "green"
	Cyan        = "aqua"
	Red         = "red"
	Pink        = "light_purple"
	Yellow      = "yellow"
	White       = "white"
)

func RGB(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func CustomColor(c color.Color) string {
	r, g, b, _ := c.RGBA()

	return RGB(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}
