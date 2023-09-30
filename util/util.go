package util

import (
	"os"
	"strings"
)

func AddDashesToUUID(uuid string) string {
	str := ""
	for i, char := range strings.Split(uuid, "") {
		str += char
		if i == 7 || i == 11 || i == 15 || i == 19 {
			str += "-"
		}
	}
	return str
}

func HasArg(arg string) bool {
	for _, s := range os.Args {
		if s == arg {
			return true
		}
	}
	return false
}

type Placeholders struct {
	PlayerName   string
	Message      string
	PlayerGroup  string
	PlayerPrefix string
	PlayerSuffix string
}

func ParsePlaceholders(str string, placeholders Placeholders) string {
	str = strings.ReplaceAll(str, "%player%", placeholders.PlayerName)
	str = strings.ReplaceAll(str, "%message%", placeholders.Message)
	str = strings.ReplaceAll(str, "%player_prefix%", placeholders.PlayerPrefix)
	str = strings.ReplaceAll(str, "%player_suffix%", placeholders.PlayerSuffix)
	str = strings.ReplaceAll(str, "%player_group%", placeholders.PlayerGroup)
	str = strings.TrimSpace(str)
	return str
}
