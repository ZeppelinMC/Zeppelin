package session

import (
	"net"

	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net/packet/login"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/player"
	"github.com/google/uuid"
)

type Session interface {
	// Username of the session
	Username() string
	// UUID of the session
	UUID() uuid.UUID
	// Properties (typically only textures)
	Properties() []login.Property

	// The player this session holds
	Player() *player.Player
	// The client name this session reports in minecraft:brand (vanilla)
	ClientName() string
	// The address of this connection
	Addr() net.Addr
	// The broadcaster this session uses
	Broadcast() *Broadcast

	// Logins the session to the server, on the standard session this starts listening to packets too
	Login() error
	// Disconnects the session from the server
	Disconnect(reason chat.TextComponent) error
	// sends a player chat message packet to the session
	PlayerChatMessage(pk play.ChatMessage, sender Session, chatType int) error
	// sends a player info update packet to the session
	PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error

	// Returns the session data for this session, and if it has any
	SessionData() (data play.PlayerSession, ok bool)
}
