package configuration

import (
	"github.com/zeppelinmc/zeppelin/net/packet"
)

const (
	//clientbound
	PacketIdFinishConfiguration = 0x03
	//serverbound
	PacketIdAcknowledgeFinishConfiguration
)

type FinishConfiguration struct {
	packet.EmptyPacket
}

func (FinishConfiguration) ID() int32 {
	return 0x03
}

type AcknowledgeFinishConfiguration = FinishConfiguration
