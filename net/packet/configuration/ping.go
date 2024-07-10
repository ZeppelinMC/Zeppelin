package configuration

import "aether/net/io"

//two-sided
const PacketIdPing = 0x05

type Ping struct {
	ID_ int32
}

func (Ping) ID() int32 {
	return 0x05
}

func (p *Ping) Encode(w io.Writer) error {
	return w.Int(p.ID_)
}

func (p *Ping) Decode(r io.Reader) error {
	return r.Int(&p.ID_)
}

type Pong = Ping
