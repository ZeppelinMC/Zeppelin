package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// serverbound
const PacketIdChatCommand = 0x04

type ChatCommand struct {
	Command string
}

func (ChatCommand) ID() int32 {
	return PacketIdChatCommand
}

func (c *ChatCommand) Encode(w io.Writer) error {
	return w.String(c.Command)
}

func (c *ChatCommand) Decode(r io.Reader) error {
	return r.String(&c.Command)
}
