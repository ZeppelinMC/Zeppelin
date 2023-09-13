package server

import (
	"encoding/hex"
	"encoding/json"
	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"net"
	"os"
)

type user struct {
	Ip string `json:"ip,omitempty"`

	UUID string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`

	Created string `json:"created,omitempty"`
	Source  string `json:"source,omitempty"`
	Expires string `json:"expires,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func loadUsers(path string) ([]user, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := json.NewDecoder(file)

	var users []user
	err = d.Decode(&users)

	return users, err
}

func WritePlayerList(path string, user []user) error {
	data, _ := json.Marshal(user)
	return os.WriteFile(path, data, 0755)
}

// ValidateConn checks if the connection is allowed to join the server,
// if not the connection is kicked with the appropriate message.
// returns true if the connection was disconnected, false otherwise.
func (srv *Server) ValidateConn(conn *minecraft.Conn) bool {
	var reason string
	if srv.IsPlayerBanned(conn.Info.UUID) {
		reason = srv.Config.Messages.Banned
	}

	ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	if srv.IsIPBanned(ip) {
		reason = srv.Config.Messages.Banned
	}

	if srv.Config.Whitelist.Enable {
		if !srv.IsWhitelisted(conn.Info.Name) {
			reason = srv.Config.Messages.NotInWhitelist
		}
	}

	if reason != "" {
		msg := chat.NewMessage(reason)
		conn.SendPacket(&packet.DisconnectLogin{Reason: msg.String()})
	}

	return reason != ""
}

func (srv *Server) IsPlayerBanned(uuid [16]byte) bool {
	suuid := hex.EncodeToString(uuid[:])
	for _, u := range srv.BannedPlayers {
		if u.UUID == suuid {
			return true
		}
	}

	return false
}

func (srv *Server) IsIPBanned(ip string) bool {
	for _, u := range srv.BannedIPs {
		if u.Ip == ip {
			return true
		}
	}

	return false
}

func (srv *Server) IsWhitelisted(name string) bool {
	for _, u := range srv.WhitelistedPlayers {
		if u.Name == name {
			return true
		}
	}

	return false
}
