package server

import (
	"encoding/json"
	"os"
)

func loadLang(data *map[string]string) error {
	file, err := os.Open("lang.json")
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(data)
}

func createLang(data map[string]string) error {
	file, err := os.Create("lang.json")
	if err != nil {
		return err
	}
	enc := json.NewEncoder(file)
	enc.SetIndent("", "\t")
	return enc.Encode(data)
}

var defaultLang = map[string]string{
	"player.join":                               "&e%player% has joined the game",
	"player.leave":                              "&e%player% has left the game",
	"disconnect.not_whitelisted":                "You are not white-listed on this server!",
	"disconnect.server_shutdown":                "Server closed",
	"disconnect.invalid_player_movement":        "Invalid move player packet received",
	"disconnect.banned":                         "You are banned from this server",
	"disconnect.banned.reason":                  "You are banned from this server.\nReason: %reason%",
	"commands.deop.success":                     "Made %player% no longer a server operator",
	"commands.gamemode.success.other":           "Set %player%'s game mode to %gamemode%",
	"commands.gamemode.success.self":            "Set own game mode to %gamemode%",
	"commands.kill.success.single":              "Killed %player%",
	"commands.op.success":                       "Made %player% a server operator",
	"commands.reload.success":                   "Reloading!",
	"commands.pardon.success":                   "Unbanned %player%",
	"commands.seed.success":                     "Seed: [&a%seed%&f]",
	"commands.teleport.success.entity.single":   "Teleported %player% to %player1%",
	"commands.teleport.success.location.single": "Teleported %player% to %x%, %y%, %z%",
}
