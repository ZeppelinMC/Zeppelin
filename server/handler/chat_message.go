package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func ChatMessagePacket(state *player.Player, pk *packet.ChatMessageServer) {
	state.HandleChat(pk)
}
