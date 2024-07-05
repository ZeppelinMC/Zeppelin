package net

import (
	"aether/net/packet"
	"aether/net/packet/handshake"
	"aether/net/packet/login"
	"aether/net/packet/status"
)

var serverboundPool = map[int]map[int32]func() packet.Packet{
	HandshakingState: {
		0x00: func() packet.Packet { return &handshake.Handshaking{} },
	},
	StatusState: {
		0x00: func() packet.Packet { return &status.StatusRequest{} },
		0x01: func() packet.Packet { return &status.Ping{} },
	},
	LoginState: {
		0x00: func() packet.Packet { return &login.LoginStart{} },
		0x01: func() packet.Packet { return &login.EncryptionResponse{} },
		0x03: func() packet.Packet { return &login.LoginAcknowledged{} },
		0x04: func() packet.Packet { return &login.CookieResponse{} },
	},
}
