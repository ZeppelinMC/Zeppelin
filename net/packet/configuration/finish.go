package configuration

import (
	"aether/net/packet"
)

type FinishConfiguration struct {
	packet.EmptyPacket
}

func (FinishConfiguration) ID() int32 {
	return 0x03
}

type AcknowledgeFinishConfiguration = FinishConfiguration
