package handlers

func PlayerSession(controller Controller, id [16]byte, pk []byte) {
	controller.SetSessionID(id, pk)
}
