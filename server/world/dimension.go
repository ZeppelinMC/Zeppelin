package world

type Dimension struct {
	typ string
	//TODO
	//chunk reader
	//chunk generator
	//chunk writer

	seed int64
}

func NewDimension(typ string) *Dimension {
	return &Dimension{typ: typ}
}

func (d *Dimension) Type() string {
	return d.typ
}

func (d *Dimension) Seed() int64 {
	return d.seed
}
