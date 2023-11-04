package permission

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

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

var cache = struct {
	players map[string]PlayerPermissions
	groups  map[string]GroupPermissions
	mu      sync.RWMutex
}{
	players: make(map[string]PlayerPermissions),
	groups:  make(map[string]GroupPermissions),
}

func Save() {
	os.MkdirAll("permissions/players", 0755)
	os.Mkdir("permissions/groups", 0755)
	for p, d := range cache.players {
		j, _ := json.Marshal(d)
		os.WriteFile(fmt.Sprintf("permissions/players/%s.json", p), j, 0755)
	}
	for g, d := range cache.groups {
		j, _ := json.Marshal(d)
		os.WriteFile(fmt.Sprintf("permissions/groups/%s.json", g), j, 0755)
	}
}

func Clear() {
	cache.mu.Lock()
	clear(cache.players)
	clear(cache.groups)
	cache.mu.Unlock()
}

func GetPlayer(playerId string) PlayerPermissions {
	cache.mu.RLock()
	if cache.players[playerId].Permissions != nil {
		cache.mu.RUnlock()
		return cache.players[playerId]
	}
	cache.mu.RUnlock()
	d, err := os.ReadFile(fmt.Sprintf("permissions/players/%s.json", playerId))
	if err != nil {
		p := PlayerPermissions{
			Group: "default",
		}
		cache.mu.Lock()
		cache.players[playerId] = p
		cache.mu.Unlock()
		return p
	}
	var data PlayerPermissions
	json.Unmarshal(d, &data)
	return data
}

func GetGroup(group string) GroupPermissions {
	cache.mu.RLock()
	if cache.groups[group].Permissions != nil {
		cache.mu.RUnlock()
		return cache.groups[group]
	}
	cache.mu.RUnlock()

	d, err := os.ReadFile(fmt.Sprintf("permissions/groups/%s.json", group))
	if err != nil {
		p := GroupPermissions{Permissions: map[string]bool{"server.chat": true}}
		cache.mu.Lock()
		cache.groups[group] = p
		cache.mu.Unlock()
		return p
	}
	var data GroupPermissions
	json.Unmarshal(d, &data)
	return data
}

func HasPermissions(p string, perms []string) bool {
	if len(perms) == 0 {
		return true
	}
	permissionsPlayer := GetPlayer(p)
	permissionsGroup := GetGroup(permissionsPlayer.Group)
	for _, perm := range perms {
		if !permissionsPlayer.Permissions[perm] && !permissionsGroup.Permissions[perm] {
			return false
		}
	}
	return true
}
