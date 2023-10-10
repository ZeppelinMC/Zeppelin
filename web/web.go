package web

import (
	"embed"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/util"
	"github.com/gorilla/websocket"
)

//go:embed pages cdn
var guifs embed.FS

var log *logger.Logger

type handler struct {
	password string
}

type conn struct {
	conn *websocket.Conn
	auth bool
}

var upgrader = websocket.Upgrader{}

var conns = make([]*conn, 0)

func (h *handler) Render(w http.ResponseWriter, name string, vars map[string]string) (int, error) {
	f, err := os.ReadFile("web/pages/" + name) //guifs.ReadFile("pages/" + name)
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

func (h *handler) HandleConn(c *conn) {
	var msg map[string]interface{}
	for {
		if c.conn.ReadJSON(&msg) != nil {
			for i, co := range conns {
				if co == c {
					conns[i] = nil
				}
			}
			return
		}
		switch msg["type"] {
		case "auth":
			pass := msg["data"]
			if pass != h.password {
				c.conn.WriteJSON(map[string]interface{}{
					"type": "error",
					"data": "Wrong password",
				})
				c.conn.Close()
				c = nil
				return
			}
			c.auth = true
			c.conn.WriteJSON(syncLog(strings.Join(messages, "\n")))
		}
	}
}

func (h *handler) Upgrade(w http.ResponseWriter, r *http.Request) {
	c, _ := upgrader.Upgrade(w, r, nil)
	co := &conn{
		conn: c,
	}
	conns = append(conns, co)
	go h.HandleConn(co)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrade := false
	for _, header := range r.Header["Upgrade"] {
		if header == "websocket" {
			upgrade = true
			break
		}
	}
	if upgrade {
		h.Upgrade(w, r)
	} else {
		uri, _ := url.ParseRequestURI(r.RequestURI)
		var code int
		if strings.HasPrefix(r.RequestURI, "/cdn") {
			file, err := os.ReadFile("web/cdn/" + strings.TrimPrefix(uri.Path, "/cdn/")) //guifs.ReadFile("cdn/" + strings.TrimPrefix(uri.Path, "/cdn/"))
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
		if strings.HasPrefix(r.RequestURI, "/api") {
			p := strings.TrimPrefix(uri.Path, "/api/")
			switch p {
			case "login":
				h.Login(uri.Query().Get("p"), w)
			}
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
}

func logMessage(m string) map[string]interface{} {
	return map[string]interface{}{
		"type": "log",
		"data": m,
	}
}

func syncLog(m string) map[string]interface{} {
	return map[string]interface{}{
		"type": "sync",
		"data": m,
	}
}

var messages []string

func LaunchWebPanel(addr string, password string, l *logger.Logger) {
	go func() {
		for d := range l.Channel() {
			msg := string(extract(json.Marshal(d)))
			messages = append(messages, msg)
			for _, c := range conns {
				if c == nil {
					continue
				}
				if c.auth {
					c.conn.WriteJSON(logMessage(msg))
				}
			}
		}
	}()
	l.EnableChannel()
	log = l
	if len(password) < 8 && !util.HasArg("-no_password_req") {
		log.Error("Failed to start web panel. Password must be at least 8 characters long. You can bypass this using -no_password_req")
		return
	}
	log.Info("Launching web panel at http://%s", addr)
	err := http.ListenAndServe(addr, &handler{password})
	if err != nil {
		log.Error("Failed to start web panel: %s", err)
	}
}

func extract[T any](t T, err error) T {
	return t
}
