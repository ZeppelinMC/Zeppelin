package status

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

type Ping struct {
	Payload int64
}

func (Ping) ID() int32 {
	return 0x01
}

func (p Ping) Encode(w encoding.Writer) error {
	return w.Long(p.Payload)
}

func (p *Ping) Decode(r encoding.Reader) error {
	return r.Long(&p.Payload)
}
