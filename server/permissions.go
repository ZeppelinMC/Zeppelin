package server

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	Permissions:
		server.command.stop - /stop command
		server.command.reload - /reload command
		server.command.op - /op command
		server.command.gamemode - /gamemode command
		server.chat - Use chat
		server.chat.colors - Use chat colors
*/

var groupCache = make(map[string]GroupPermissions)
var playerCache = make(map[string]PlayerPermissions)

type PlayerPermissions struct {
	Group       string          `json:"group"`
	Permissions map[string]bool `json:"permissions"`
}

type GroupPermissions struct {
	DisplayName string          `json:"display_name"`
	Prefix      string          `json:"prefix"`
	Suffix      string          `json:"suffix"`
	Permissions map[string]bool `json:"permissions"`
}

func getPlayer(playerId string) PlayerPermissions {
	if playerCache[playerId].Permissions != nil {
		return playerCache[playerId]
	}
	d, err := os.ReadFile(fmt.Sprintf("permissions/players/%s.json", playerId))
	if err != nil {
		os.WriteFile(fmt.Sprintf("permissions/players/%s.json", playerId), []byte(`{"group":"default"}`), 0755)
		return PlayerPermissions{}
	}
	var data PlayerPermissions
	json.Unmarshal(d, &data)
	return data
}

func getGroup(group string) GroupPermissions {
	if groupCache[group].Permissions != nil {
		return groupCache[group]
	}
	d, err := os.ReadFile(fmt.Sprintf("permissions/groups/%s.json", group))
	if err != nil {
		return GroupPermissions{}
	}
	var data GroupPermissions
	json.Unmarshal(d, &data)
	return data
}

func (p *PlayerController) HasPermissions(perms []string) bool {
	if len(perms) == 0 {
		return true
	}
	if p.player.Operator {
		return true
	}
	permissionsPlayer := getPlayer(p.UUID)
	permissionsGroup := getGroup(permissionsPlayer.Group)
	for _, perm := range perms {
		if !permissionsPlayer.Permissions[perm] && !permissionsGroup.Permissions[perm] {
			return false
		}
	}
	return true
}
