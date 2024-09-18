package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdDisguisedChatMessage = 0x1E

type DisguisedChatMessage struct {
	Message text.TextComponent

	ChatType   int32
	SenderName text.TextComponent

	TargetName *text.TextComponent
}

func (DisguisedChatMessage) ID() int32 {
	return PacketIdDisguisedChatMessage
}

func (p *DisguisedChatMessage) Encode(w encoding.Writer) error {
	if err := w.TextComponent(p.Message); err != nil {
		return err
	}

	if err := w.VarInt(p.ChatType); err != nil {
		return err
	}
	if err := w.TextComponent(p.SenderName); err != nil {
		return err
	}

	if err := w.Bool(p.TargetName != nil); err != nil {
		return err
	}
	if p.TargetName != nil {
		if err := w.TextComponent(*p.TargetName); err != nil {
			return err
		}
	}

	return nil
}

func (p *DisguisedChatMessage) Decode(r encoding.Reader) error {
	if err := r.TextComponent(&p.Message); err != nil {
		return err
	}

	if _, err := r.VarInt(&p.ChatType); err != nil {
		return err
	}
	if err := r.TextComponent(&p.SenderName); err != nil {
		return err
	}

	var hasTgt bool
	if err := r.Bool(&hasTgt); err != nil {
		return err
	}
	if hasTgt {
		if err := r.TextComponent(p.TargetName); err != nil {
			return err
		}
	}

	return nil
}
