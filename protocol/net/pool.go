package net

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/handshake"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/login"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/status"
)

var ServerboundPool = map[int32]map[int32]func() packet.Decodeable{
	HandshakingState: {
		0x00: func() packet.Decodeable { return &handshake.Handshaking{} },
	},
	StatusState: {
		0x00: func() packet.Decodeable { return &status.StatusRequest{} },
		0x01: func() packet.Decodeable { return &status.Ping{} },
	},
	LoginState: {
		0x00: func() packet.Decodeable { return &login.LoginStart{} },
		0x02: func() packet.Decodeable { return &login.LoginPluginResponse{} },
		0x01: func() packet.Decodeable { return &login.EncryptionResponse{} },
		0x03: func() packet.Decodeable { return &login.LoginAcknowledged{} },
		0x04: func() packet.Decodeable { return &login.CookieResponse{} },
	},
	ConfigurationState: {
		0x00: func() packet.Decodeable { return &configuration.ClientInformation{} },
		0x01: func() packet.Decodeable { return &configuration.CookieResponse{} },
		0x02: func() packet.Decodeable { return &configuration.ServerboundPluginMessage{} },
		0x03: func() packet.Decodeable { return &configuration.AcknowledgeFinishConfiguration{} },
		0x04: func() packet.Decodeable { return &configuration.KeepAlive{} },
		0x05: func() packet.Decodeable { return &configuration.Pong{} },
	},
	PlayState: {
		0x00: func() packet.Decodeable { return &play.ConfirmTeleportation{} },
		0x04: func() packet.Decodeable { return &play.ChatCommand{} },
		0x05: func() packet.Decodeable { return &play.SignedChatCommand{} },
		0x06: func() packet.Decodeable { return &play.ChatMessage{} },
		0x08: func() packet.Decodeable { return &play.ChunkBatchReceived{} },
		0x07: func() packet.Decodeable { return &play.PlayerSession{} },
		0x0A: func() packet.Decodeable { return &play.ClientInformation{} },
		0x0E: func() packet.Decodeable { return &play.ClickContainer{} },
		0x0F: func() packet.Decodeable { return &play.CloseContainer{} },
		0x12: func() packet.Decodeable { return &play.ServerboundPluginMessage{} },
		0x16: func() packet.Decodeable { return &play.Interact{} },
		0x18: func() packet.Decodeable { return &play.ServerboundKeepAlive{} },
		0x1A: func() packet.Decodeable { return &play.SetPlayerPosition{} },
		0x1B: func() packet.Decodeable { return &play.SetPlayerPositionAndRotation{} },
		0x1C: func() packet.Decodeable { return &play.SetPlayerRotation{} },
		0x1D: func() packet.Decodeable { return &play.SetPlayerOnGround{} },
		0x23: func() packet.Decodeable { return &play.PlayerAbilitiesServerbound{} },
		0x25: func() packet.Decodeable { return &play.PlayerCommand{} },
		0x2F: func() packet.Decodeable { return &play.SetHeldItemServerbound{} },
		0x32: func() packet.Decodeable { return &play.SetCreativeModeSlot{} },
		0x36: func() packet.Decodeable { return &play.SwingArm{} },
		0x38: func() packet.Decodeable { return &play.UseItemOn{} },
	},
}
