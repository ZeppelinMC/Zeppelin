package world

type Dimension struct {
	typ string
	//TODO
	//chunk reader
	//chunk gnerator
	//chunk writer
}

func NewDimension(typ string) *Dimension {
	return &Dimension{typ: typ}
}

func (d *Dimension) Type() string {
	return d.typ
}
