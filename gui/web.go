package gui

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"strings"

	"github.com/dynamitemc/dynamite/util"
)

type logger interface {
	Error(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
}

//go:embed pages cdn
var guifs embed.FS

var log logger

type handler struct{}

func (handler) Render(w http.ResponseWriter, name string, vars map[string]string) (int, error) {
	f, err := guifs.ReadFile("gui/pages/" + name)
	if err != nil {
		return 0, err
	}
	w.Header().Set("Content-Type", "text/html")
	file := string(f)
	for k, v := range vars {
		file = strings.ReplaceAll(file, "{{"+k+"}}", v)
	}
	return io.WriteString(w, file)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri, _ := url.ParseRequestURI(r.RequestURI)
	var code int
	if strings.HasPrefix(r.RequestURI, "/cdn") {
		file, err := guifs.ReadFile("gui/cdn/" + strings.TrimPrefix(uri.Path, "/cdn/"))
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				io.WriteString(w, "Unknown file!")
				return
			}
			io.WriteString(w, "Failed to open file!")
			return
		}
		io.WriteString(w, string(file))
		return
	}
	switch uri.Path {
	case "/":
		{
			code = http.StatusOK
			h.Render(w, "panel.html", nil)
		}
	case "/login":
		{
			code = http.StatusOK
			h.Render(w, "login.html", nil)
		}
	default:
		{
			code = http.StatusNotFound
			h.Render(w, "notfound.html", nil)
		}
	}
	log.Debug("[WEB] [%s] visited %s | Code %d", r.RemoteAddr, r.RequestURI, code)
}

func LaunchWebPanel(addr string, password string, l logger) {
	log = l
	if len(password) < 8 && !util.HasArg("-no_password_req") {
		log.Error("Failed to start web panel. Password must be at least 8 characters long. You can bypass this using -no_password_req")
		return
	}
	log.Info("Launching web panel at http://%s", addr)
	err := http.ListenAndServe(addr, handler{})
	if err != nil {
		log.Error("Failed to start web panel: %s", err)
	}
}
