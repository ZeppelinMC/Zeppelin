package status

import "github.com/dynamitemc/aether/net/io"

type Ping struct {
	Payload int64
}

func (Ping) ID() int32 {
	return 0x01
}

func (p Ping) Encode(w io.Writer) error {
	return w.Long(p.Payload)
}

func (p *Ping) Decode(r io.Reader) error {
	return r.Long(&p.Payload)
}
