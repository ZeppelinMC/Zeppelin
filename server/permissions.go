package server

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func saveCache() {
	os.MkdirAll("permissions/players", 0755)
	os.Mkdir("permissions/groups", 0755)
	for p, d := range playerCache {
		j, _ := json.Marshal(d)
		os.WriteFile(fmt.Sprintf("permissions/players/%s.json", p), j, 0755)
	}
	for g, d := range groupCache {
		j, _ := json.Marshal(d)
		os.WriteFile(fmt.Sprintf("permissions/groups/%s.json", g), j, 0755)
	}
}

func clearCache() {
	clear(playerCache)
	clear(groupCache)
}

func getPlayer(playerId string) PlayerPermissions {
	if playerCache[playerId].Permissions != nil {
		return playerCache[playerId]
	}
	d, err := os.ReadFile(fmt.Sprintf("permissions/players/%s.json", playerId))
	if err != nil {
		p := PlayerPermissions{
			Group: "default",
		}
		playerCache[playerId] = p
		return p
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
		p := GroupPermissions{Permissions: map[string]bool{"server.chat": true}}
		groupCache[group] = p
		return p
	}
	var data GroupPermissions
	json.Unmarshal(d, &data)
	return data
}

func (p *PlayerController) HasPermissions(perms []string) bool {
	if len(perms) == 0 {
		return true
	}
	if p.player.Operator() {
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
