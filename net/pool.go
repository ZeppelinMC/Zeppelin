package net

import (
	"aether/net/packet"
	"aether/net/packet/configuration"
	"aether/net/packet/handshake"
	"aether/net/packet/login"
	"aether/net/packet/play"
	"aether/net/packet/status"
)

var serverboundPool = map[int32]map[int32]func() packet.Packet{
	HandshakingState: {
		0x00: func() packet.Packet { return &handshake.Handshaking{} },
	},
	StatusState: {
		0x00: func() packet.Packet { return &status.StatusRequest{} },
		0x01: func() packet.Packet { return &status.Ping{} },
	},
	LoginState: {
		0x00: func() packet.Packet { return &login.LoginStart{} },
		0x02: func() packet.Packet { return &login.LoginPluginResponse{} },
		0x01: func() packet.Packet { return &login.EncryptionResponse{} },
		0x03: func() packet.Packet { return &login.LoginAcknowledged{} },
		0x04: func() packet.Packet { return &login.CookieResponse{} },
	},
	ConfigurationState: {
		0x00: func() packet.Packet { return &configuration.ClientInformation{} },
		0x01: func() packet.Packet { return &configuration.CookieResponse{} },
		0x02: func() packet.Packet { return &configuration.ServerboundPluginMessage{} },
		0x03: func() packet.Packet { return &configuration.AcknowledgeFinishConfiguration{} },
		0x04: func() packet.Packet { return &configuration.KeepAlive{} },
		0x05: func() packet.Packet { return &configuration.Pong{} },
	},
	PlayState: {
		0x0A: func() packet.Packet { return &play.ClientInformation{} },
		0x12: func() packet.Packet { return &play.ServerboundPluginMessage{} },
	},
}
