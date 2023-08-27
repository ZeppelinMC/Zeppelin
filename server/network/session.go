package network

import "github.com/aimjel/minecraft"

type Session struct {
	Conn *minecraft.Conn
}

func NewSession(c *minecraft.Conn) *Session {
	return &Session{Conn: c}
}
