package gui

import (
	"io"
	"net/http"

	"github.com/dynamitemc/dynamite/util"
)

type logger interface {
	Error(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

var log logger

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var code int
	switch r.RequestURI {
	case "/":
		{
			io.WriteString(w, "hello there!")
		}
	}
	log.Debug("[GUI] [%s] visited %s | Code %d", r.RemoteAddr, r.RequestURI, code)
}

func LaunchGUI(addr string, password string, l logger) {
	log = l
	if len(password) < 8 && !util.HasArg("-no_password_req") {
		log.Error("Failed to start http gui panel. Password must be at least 8 characters long. You can bypass this using -no_password_req")
		return
	}
	log.Info("Launching http gui panel at http://%s", addr)
	err := http.ListenAndServe(addr, handler{})
	if err != nil {
		log.Error("Failed to start http gui panel: %s", err)
	}
}
