package play

import "github.com/zeppelinmc/zeppelin/protocol/net/packet"

// clientbound
const PacketIdBundleDelimiter = 0x00

type BundleDelimiter struct{ packet.EmptyPacket }

func (BundleDelimiter) ID() int32 {
	return PacketIdBundleDelimiter
}
