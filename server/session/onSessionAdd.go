package session

import "github.com/zeppelinmc/zeppelin/protocol/text"

func onSessionAdd(s Session) {
	s.Broadcast().SystemChatMessage(text.Unmarshalf('&', "&e%s joined the game", s.Username()))
}
