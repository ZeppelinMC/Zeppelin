package handler

import (
	"fmt"
	"runtime"

	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	"github.com/zeppelinmc/zeppelin/text"
)

func init() {
	std.RegisterHandler(net.PlayState, play.PacketIdChatMessage, handleChatMessage)
}

func handleChatMessage(s *std.StandardSession, pk packet.Packet) {
	if cm, ok := pk.(*play.ChatMessage); ok {
		if len(cm.Message) > 256 {
			s.Disconnect(text.TextComponent{Text: "Chat message over 256 characters is not allowed"})
			return
		}
		if cm.Message == "#stat" {
			var stats runtime.MemStats
			runtime.ReadMemStats(&stats)

			s.SystemMessage(text.TextComponent{
				Text: fmt.Sprintf("Server stats: \n\n Alloc: %dMiB, Total Alloc: %dMiB\n\n Why are you using illegal commands though?", stats.Alloc/1024/1024, stats.TotalAlloc/1024/1024),
			})
		} else {
			switch s.Config().Chat.ChatMode {
			case "secure":
				i, p := s.SecureChatData()
				s.Broadcast().SecureChatMessage(s, *cm, i, p)
				s.AppendMessage(cm.Signature)
			case "disguised":
				comp := text.TextComponent{Text: cm.Message}
				if s.Config().Chat.Colors {
					comp = text.Unmarshal(cm.Message, rune(s.Config().Chat.Formatter[0]))
				}
				s.Broadcast().DisguisedChatMessage(s, comp)
			default: // off
				return
			}
		}
	}
}
