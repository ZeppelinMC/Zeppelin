package server

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"slices"
	"time"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	"github.com/google/uuid"
)

type user struct {
	Ip string `json:"ip,omitempty"`

	UUID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`

	Created string `json:"created,omitempty"`
	Source  string `json:"source,omitempty"`
	Expires string `json:"expires,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func loadUsers(path string) ([]user, error) {
	file, err := os.Open(path)
	if err != nil {
		os.WriteFile(path, []byte("[]"), 0755)
		file, err = os.Open(path)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
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
		if !srv.IsWhitelisted(conn.Info.UUID) {
			reason = srv.Config.Messages.NotInWhitelist
		}
	}

	if reason != "" {
		conn.SendPacket(&packet.DisconnectPlay{DisconnectLogin: packet.DisconnectLogin{Reason: reason}})
	}

	return reason != ""
}

func (srv *Server) IsPlayerBanned(u [16]byte) bool {
	suuid, _ := uuid.FromBytes(u[:])
	for _, u := range srv.BannedPlayers {
		if u.UUID == suuid.String() {
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

func (srv *Server) IsWhitelisted(u [16]byte) bool {
	suuid, _ := uuid.FromBytes(u[:])
	for _, u := range srv.WhitelistedPlayers {
		if u.UUID == suuid.String() {
			return true
		}
	}

	return false
}

func (srv *Server) IsOperator(uuid [16]byte) bool {
	suuid := hex.EncodeToString(uuid[:])
	for _, u := range srv.WhitelistedPlayers {
		if u.UUID == suuid {
			return true
		}
	}

	return false
}

func (srv *Server) Ban(p *PlayerController, reason string) {
	t, _ := time.Now().MarshalJSON()
	srv.BannedPlayers = append(srv.BannedPlayers, user{
		UUID:    p.UUID,
		Name:    p.Name(),
		Created: string(t),
		Reason:  reason,
	})
}

func (srv *Server) Unban(p *PlayerController) {
	for i, b := range srv.BannedPlayers {
		if b.UUID == p.UUID {
			srv.BannedPlayers = slices.Delete(srv.BannedPlayers, i, i+1)
			return
		}
	}
}

func (srv *Server) MakeOperator(p *PlayerController) {
	p.player.SetOperator(true)
	srv.Operators = append(srv.Operators, user{
		UUID: p.UUID,
		Name: p.Name(),
	})
}

func (srv *Server) MakeNotOperator(p *PlayerController) {
	p.player.SetOperator(false)
	for i, op := range srv.Operators {
		if op.UUID == p.UUID {
			srv.Operators = slices.Delete(srv.Operators, i, i+1)
			return
		}
	}
}

func (srv *Server) AddToWhitelist(p *PlayerController) {
	srv.WhitelistedPlayers = append(srv.WhitelistedPlayers, user{
		UUID: p.UUID,
		Name: p.Name(),
	})
}

func (srv *Server) RemoveFromWhitelist(p *PlayerController) {
	for i, w := range srv.WhitelistedPlayers {
		if w.UUID == p.UUID {
			srv.WhitelistedPlayers = slices.Delete(srv.WhitelistedPlayers, i, i+1)
			return
		}
	}
}
