package handlers

import "github.com/aimjel/minecraft/packet"

func ChatMessagePacket(controller Controller, pk *packet.ChatMessageServer) {
	controller.Chat(pk)
}
