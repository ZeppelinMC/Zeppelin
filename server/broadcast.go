package server

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/player"
)

func (srv Server) GlobalBroadcast(pk packet.Packet) {
	srv.Lock()
	defer srv.Unlock()
	for _, player := range srv.Players {
		player.Session.Conn.SendPacket(pk)
	}
}

func (srv Server) PlayerlistUpdate() {
	var players []player.Info
	srv.Lock()
	for _, p := range srv.Players {
		info := *p.Session.Conn.Info
		info.Listed = true
		players = append(players, info)
	}
	srv.Unlock()
	srv.GlobalBroadcast(&packet.PlayerInfoUpdate{
		Actions: 0x01 | 0x08,
		Players: players,
	})
}

func (srv Server) PlayerlistRemove(players ...[16]byte) {
	srv.GlobalBroadcast(&packet.PlayerInfoRemove{UUIDS: players})
}
