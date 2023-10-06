package handlers

func ChatMessagePacket(controller Controller, content string) {
	controller.Chat(content)
}
