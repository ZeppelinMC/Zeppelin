package std

import (
	"slices"

	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/text"
)

func (session *StandardSession) PlayerChatMessage(
	pk play.ChatMessage, sender session.Session, chatType string,
	index int32, prevMsgs []play.PreviousMessage,
) error {
	chatTypeIndex := slices.Index(session.registryIndexes["minecraft:chat_type"], chatType)

	return session.conn.WritePacket(&play.PlayerChatMessage{
		Sender:              sender.UUID(),
		Index:               index,
		HasMessageSignature: pk.HasSignature,
		MessageSignature:    pk.Signature,
		Message:             pk.Message,
		Timestamp:           pk.Timestamp,
		Salt:                pk.Salt,

		PreviousMessages: session.previousMessages,

		ChatType:   int32(chatTypeIndex + 1),
		SenderName: text.Unmarshal(sender.Username(), rune(session.config.Chat.Formatter[0])),
	})
}

func (session *StandardSession) DisguisedChatMessage(content text.TextComponent, sender session.Session, chatType string) error {
	chatTypeIndex := slices.Index(session.registryIndexes["minecraft:chat_type"], chatType)

	return session.conn.WritePacket(&play.DisguisedChatMessage{
		Message: content,

		ChatType:   int32(chatTypeIndex + 1),
		SenderName: text.Unmarshal(sender.Username(), rune(session.config.Chat.Formatter[0])),
	})
}

func (session *StandardSession) AppendMessage(sig [256]byte) {
	session.prev_msgs_mu.Lock()
	defer session.prev_msgs_mu.Unlock()
	session.bumpChatIndex()
	session.previousMessages = append(session.previousMessages, play.PreviousMessage{MessageID: -1, Signature: &sig})

	if len(session.previousMessages) > 20 {
		session.previousMessages = session.previousMessages[1:21]
	}
}

func (session *StandardSession) bumpChatIndex() {
	session.ChatIndex.Set(session.ChatIndex.Get() + 1)
}

func (session *StandardSession) SecureChatData() (index int32, prevMsgs []play.PreviousMessage) {
	session.prev_msgs_mu.Lock()
	defer session.prev_msgs_mu.Unlock()

	return session.ChatIndex.Get(), session.previousMessages
}

func (session *StandardSession) SystemMessage(msg text.TextComponent) error {
	return session.conn.WritePacket(&play.SystemChatMessage{Content: msg})
}
