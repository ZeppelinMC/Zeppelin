package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdSystemChatMessage = 0x6C

type SystemChatMessage struct {
	Content text.TextComponent
	Overlay bool
}

func (SystemChatMessage) ID() int32 {
	return PacketIdSystemChatMessage
}

func (s *SystemChatMessage) Encode(w encoding.Writer) error {
	if err := w.TextComponent(s.Content); err != nil {
		return err
	}
	return w.Bool(s.Overlay)
}

func (s *SystemChatMessage) Decode(r encoding.Reader) error {
	if err := r.TextComponent(&s.Content); err != nil {
		return err
	}
	return r.Bool(&s.Overlay)
}
