package server

import (
	"fmt"
	"slices"
	"strings"

	"github.com/aimjel/minecraft/chat"
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

	net := chat.NewMessage(prefix + p.Name() + suffix).WithSuggestCommandClickEvent(fmt.Sprintf("/msg %s", p.Name()))

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
		if p.Server.Config.Chat.Format == "" {
			p.Server.mu.RLock()
			defer p.Server.mu.RUnlock()
			for _, pl := range p.Server.players {
				if pl.clientInfo.ChatMode != 0 {
					continue
				}
				pl.SendPacket(&packet.DisguisedChatMessage{
					Message:      chat.NewMessage(pk.Message),
					ChatTypeName: net,
				})
			}
		} else {
			msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
				"player":        p.Name(),
				"player_prefix": prefix,
				"player_suffix": suffix,
				"message":       pk.Message,
			})

			p.Server.GlobalMessage(msg)
		}
	} else {
		p.Server.mu.RLock()
		defer p.Server.mu.RUnlock()
		for _, pl := range p.Server.players {
			if pl.clientInfo.ChatMode != 0 {
				continue
			}
			pl.mu.Lock()
			defer pl.mu.Unlock()
			pl.SendPacket(&packet.PlayerChatMessage{
				Sender:           p.conn.UUID(),
				MessageSignature: pk.Signature,
				Index:            pl.index,
				Message:          pk.Message,
				Timestamp:        pk.Timestamp,
				Salt:             pk.Salt,
				NetworkName:      net,
				PreviousMessages: pl.acknowledgedMessages,
			})
			if len(pl.acknowledgedMessages) != 20 {
				pl.acknowledgedMessages = append([]packet.PreviousMessage{
					{
						MessageID: pl.index,
						Signature: pk.Signature,
					},
				}, pl.acknowledgedMessages...)
			}
			pl.index++
		}
	}
}

func (p *Session) Whisper(pl *Session, msg string, timestamp, salt int64, sig []byte) {
	prefix, suffix := p.GetPrefixSuffix()
	prefix1, suffix1 := pl.GetPrefixSuffix()
	tgt := chat.NewMessage(prefix1 + pl.Name() + suffix1)
	p.SendPacket(&packet.PlayerChatMessage{
		Sender:  p.conn.UUID(),
		Message: msg,
		//MessageSignature:  sig,
		Salt:              salt,
		Timestamp:         timestamp,
		ChatType:          3,
		NetworkName:       chat.NewMessage(prefix + p.Name() + suffix),
		NetworkTargetName: &tgt,
	})
}
