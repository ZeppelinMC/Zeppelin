package session

import (
	"github.com/dynamitemc/aether/net/packet/play"
)

func (b *Broadcast) ChatMessage(session Session, pk play.ChatMessage) {
	b.sessions_mu.Lock()
	defer b.sessions_mu.Unlock()

	for _, ses := range b.sessions {
		ses.PlayerChatMessage(pk, session, 0)
	}
}
