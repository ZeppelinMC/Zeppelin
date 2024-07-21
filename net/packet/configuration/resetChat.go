package configuration

import "github.com/zeppelinmc/zeppelin/net/packet"

//clientbound
const PacketIdResetChat = 0x06

type ResetChat struct{ packet.EmptyPacket }

func (ResetChat) ID() int32 {
	return 0x06
}
