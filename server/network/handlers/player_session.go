package handlers

func PlayerSession(controller Controller, id [16]byte, pk, ks []byte, expires int64) {
	controller.SetSessionID(id, pk, ks, expires)
}
