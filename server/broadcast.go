package server

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/player"
)

func (server Server) GlobalBroadcast(pk packet.Packet) {
	server.Lock()
	defer server.Unlock()
	for _, player := range server.Players {
		player.Session.Conn.SendPacket(pk)
	}
}

func (server Server) PlayerlistUpdate() {
	var players []player.Info
	server.Lock()
	for _, player := range server.Players {
		info := *player.Session.Conn.Info
		info.Listed = true
		players = append(players, info)
	}
	server.Unlock()
	server.GlobalBroadcast(&packet.PlayerInfoUpdate{
		Actions: 0x01 | 0x08,
		Players: players,
	})
}

func (server Server) PlayerlistRemove(players ...[16]byte) {
	server.GlobalBroadcast(&packet.PlayerInfoRemove{UUIDS: players})
}
