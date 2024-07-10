package play

import "aether/net/io"

type ServerboundKeepAlive struct {
	KeepAliveID int64
}

func (ServerboundKeepAlive) ID() int32 {
	return 0x18
}

func (k *ServerboundKeepAlive) Encode(w io.Writer) error {
	return w.Long(k.KeepAliveID)
}

func (k *ServerboundKeepAlive) Decode(r io.Reader) error {
	return r.Long(&k.KeepAliveID)
}

type ClientboundKeepAlive struct {
	KeepAliveID int64
}

func (ClientboundKeepAlive) ID() int32 {
	return 0x26
}

func (k *ClientboundKeepAlive) Encode(w io.Writer) error {
	return w.Long(k.KeepAliveID)
}

func (k *ClientboundKeepAlive) Decode(r io.Reader) error {
	return r.Long(&k.KeepAliveID)
}
