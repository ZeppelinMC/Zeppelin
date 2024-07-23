package session

import (
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/text"
)

func (b *Broadcast) SecureChatMessage(session Session, pk play.ChatMessage, index int32, prevMsgs []play.PreviousMessage) {
	b.console.PlayerChatMessage(pk, session, "minecraft:chat", 0, nil)
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, ses := range b.sessions {
		ses.PlayerChatMessage(pk, session, "minecraft:chat", index, prevMsgs)
	}
}

func (b *Broadcast) DisguisedChatMessage(session Session, content text.TextComponent) {
	b.console.DisguisedChatMessage(content, session, "minecraft:chat")
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, ses := range b.sessions {
		ses.DisguisedChatMessage(content, session, "minecraft:chat")
	}
}
