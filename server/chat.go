package server

import (
	"fmt"
	"slices"
	"strings"

	"github.com/aimjel/minecraft/packet"
)

func (p *Session) Chat(pk *packet.ChatMessageServer) {
	if !p.Server.Config.Chat.Enable {
		return
	}
	if !p.HasPermissions([]string{"server.chat"}) {
		return
	}

	prefix, suffix := p.GetPrefixSuffix()

	if !p.Server.Config.Chat.Secure {
		if !p.Server.Config.Chat.Enable || !p.HasPermissions([]string{"server.chat.colors"}) {
			// strip colors
			sp := strings.Split(pk.Message, "")
			for i, c := range sp {
				if c == "&" {
					if sp[i+1] != " " {
						sp = slices.Delete(sp, i, i+2)
					}
				}
			}
			pk.Message = strings.Join(sp, "")
		}
		msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
			"player":        p.Name(),
			"player_prefix": prefix,
			"player_suffix": suffix,
			"message":       pk.Message,
		})

		p.Server.GlobalMessage(msg, p)
	} else {
		p.Server.GlobalBroadcast(&packet.PlayerChatMessage{
			Sender:           p.conn.UUID(),
			MessageSignature: pk.Signature,
			Message:          pk.Message,
			Timestamp:        pk.Timestamp,
			Salt:             pk.Salt,
			NetworkName:      prefix + p.Name() + suffix,
		})
	}
}

func (p *Session) Whisper(pl *Session, msg string, timestamp, salt int64, sig []byte) {
	prefix, suffix := p.GetPrefixSuffix()
	prefix1, suffix1 := pl.GetPrefixSuffix()
	fmt.Println(msg)
	p.SendPacket(&packet.PlayerChatMessage{
		Sender:  p.conn.UUID(),
		Message: msg,
		//MessageSignature:  sig,
		Salt:              salt,
		Timestamp:         timestamp,
		ChatType:          3,
		NetworkName:       prefix + p.Name() + suffix,
		NetworkTargetName: prefix1 + pl.Name() + suffix1,
	})
}
