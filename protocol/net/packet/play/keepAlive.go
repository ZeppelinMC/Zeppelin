package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdServerboundKeepAlive = 0x18

type ServerboundKeepAlive struct {
	KeepAliveID int64
}

func (ServerboundKeepAlive) ID() int32 {
	return 0x18
}

func (k *ServerboundKeepAlive) Encode(w encoding.Writer) error {
	return w.Long(k.KeepAliveID)
}

func (k *ServerboundKeepAlive) Decode(r encoding.Reader) error {
	return r.Long(&k.KeepAliveID)
}

// clientbound
const PacketIdClientboundKeepAlive = 0x26

type ClientboundKeepAlive struct {
	KeepAliveID int64
}

func (ClientboundKeepAlive) ID() int32 {
	return 0x26
}

func (k *ClientboundKeepAlive) Encode(w encoding.Writer) error {
	return w.Long(k.KeepAliveID)
}

func (k *ClientboundKeepAlive) Decode(r encoding.Reader) error {
	return r.Long(&k.KeepAliveID)
}
