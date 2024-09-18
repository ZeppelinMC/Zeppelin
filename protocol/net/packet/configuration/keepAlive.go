package configuration

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// two-sided
const PacketIdKeepAlive = 0x04

type KeepAlive struct {
	KeepAliveID int64
}

func (KeepAlive) ID() int32 {
	return 0x04
}

func (k *KeepAlive) Encode(w encoding.Writer) error {
	return w.Long(k.KeepAliveID)
}

func (k *KeepAlive) Decode(r encoding.Reader) error {
	return r.Long(&k.KeepAliveID)
}
