package web

import (
	"net/http"
)

func (h *handler) Login(password string, res http.ResponseWriter) {
	if password != h.password {
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Wrong password"))
		return
	}
	res.WriteHeader(http.StatusAccepted)
	res.Write(nil)
}
