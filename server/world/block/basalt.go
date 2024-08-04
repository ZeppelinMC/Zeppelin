package block

type Basalt struct {
	Axis string
}

func (b Basalt) Encode() (string, BlockProperties) {
	return "minecraft:basalt", BlockProperties{
		"axis": b.Axis,
	}
}

func (b Basalt) New(props BlockProperties) Block {
	return Basalt{
		Axis: props["axis"],
	}
}