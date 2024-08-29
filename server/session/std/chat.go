package std

import (
	"slices"

	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/session"
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

		PreviousMessages: prevMsgs,

		ChatType:   int32(chatTypeIndex + 1),
		SenderName: text.Unmarshal(sender.Username(), session.config.ChatFormatter.Rune()),
	})
}

func (session *StandardSession) DisguisedChatMessage(content text.TextComponent, sender session.Session, chatType string) error {
	chatTypeIndex := slices.Index(session.registryIndexes["minecraft:chat_type"], chatType)

	return session.conn.WritePacket(&play.DisguisedChatMessage{
		Message: content,

		ChatType:   int32(chatTypeIndex + 1),
		SenderName: text.Unmarshal(sender.Username(), session.config.ChatFormatter.Rune()),
	})
}

func (session *StandardSession) SystemMessage(msg text.TextComponent) error {
	return session.conn.WritePacket(&play.SystemChatMessage{Content: msg})
}
