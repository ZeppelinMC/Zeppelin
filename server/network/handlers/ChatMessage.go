package handlers

func ChatMessagePacket(controller controller, content string) {
	controller.Chat(content)
}
