package handler

import (
	"fmt"

	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatMessage, handleChatMessage)
}

func handleChatMessage(s *std.StandardSession, pk packet.Packet) {
	if cm, ok := pk.(*play.ChatMessage); ok {
		if len(cm.Message) > 256 {
			s.Disconnect(chat.TextComponent{Text: "Chat message over 256 characters is not allowed"})
			return
		}
		fmt.Println(cm.HasSignature)
		s.Broadcast().ChatMessage(s, *cm)
	}
}
