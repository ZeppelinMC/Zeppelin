package server

import (
	"encoding/json"
	"os"

	"github.com/dynamitemc/dynamite/util"
)

func LoadPlayerList(path string) []util.Player {
	list := []util.Player{}

	file, err := os.Open(path)
	if err != nil {
		file.Close()
		file, _ := os.Create(path)
		e := json.NewEncoder(file)
		e.Encode(&list)
		return list
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&list); err != nil {
		return nil
	}

	return list
}

func WritePlayerList(path string, player util.Player) []util.Player {
	list := []util.Player{}

	b, err := os.ReadFile(path)
	if err != nil {
		list = append(list, player)
		data, _ := json.Marshal(list)
		os.WriteFile(path, data, 0755)
	}
	json.Unmarshal(b, &list)
	list = append(list, player)
	data, _ := json.Marshal(list)
	os.WriteFile(path, data, 0755)
	return list
}

func LoadIPBans() []string {
	list := []string{}

	file, err := os.Open("banned_ips.json")
	if err != nil {
		file.Close()
		file, _ := os.Create("banned_ips.json")
		e := json.NewEncoder(file)
		e.Encode(&list)
		return list
	}
	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&list); err != nil {
		return nil
	}

	return list
}

/*

0: User is valid
1: User is not in whitelist
2: User is banned
3: Server is full
4: User is already playing on another client

*/

const (
	CONNECTION_VALID = iota
	CONNECTION_PLAYER_NOT_IN_WHITELIST
	CONNECTION_PLAYER_BANNED
	CONNECTION_SERVER_FULL
	CONNECTION_PLAYER_ALREADY_PLAYING
)

func (srv *Server) ValidatePlayer(name string, id string, ip string) int {
	for _, player := range srv.BannedPlayers {
		if player.UUID == id {
			return CONNECTION_PLAYER_BANNED
		}
	}
	for _, i := range srv.BannedIPs {
		if i == ip {
			return CONNECTION_PLAYER_BANNED
		}
	}
	if srv.Config.Whitelist.Enable {
		d := false
		for _, player := range srv.WhitelistedPlayers {
			if player.UUID == id {
				d = true
				break
			}
		}
		if !d {
			return CONNECTION_PLAYER_NOT_IN_WHITELIST
		}
	}
	if srv.Players[id] != nil {
		return CONNECTION_PLAYER_ALREADY_PLAYING
	}
	if srv.Config.MaxPlayers == -1 {
		return CONNECTION_VALID
	}
	if len(srv.Players) >= srv.Config.MaxPlayers {
		return CONNECTION_SERVER_FULL
	}
	return CONNECTION_VALID
}
