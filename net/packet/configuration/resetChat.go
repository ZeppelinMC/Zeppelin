package configuration

import "aether/net/packet"

type ResetChat struct{ packet.EmptyPacket }

func (ResetChat) ID() int32 {
	return 0x06
}
