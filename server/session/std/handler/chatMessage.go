package handler

import (
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/session/std"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatMessage, handleChatMessage)
}

func handleChatMessage(s *std.StandardSession, pk packet.Decodeable) {
	if cm, ok := pk.(*play.ChatMessage); ok {
		if len(cm.Message) > 256 {
			s.Disconnect(text.TextComponent{Text: "Chat message over 256 characters is not allowed"})
			return
		}
		cfg := s.Config()

		if !cfg.EnableChat {
			return
		}
		if cfg.EnforceSecureProfile {
			i := s.ChatIndex.Get()
			s.Broadcast().SecureChatMessage(s, *cm, i)
			s.ChatIndex.Set(i + 1)
			return
		}
		if cfg.SystemChatFormat == "" {
			comp := text.TextComponent{Text: cm.Message}
			//if s.Config().Chat.Colors {
			//	comp = text.Unmarshal(cm.Message, s.Config().Chat.Formatter.Rune())
			//}
			s.Broadcast().DisguisedChatMessage(s, comp)
		}

	}
}
