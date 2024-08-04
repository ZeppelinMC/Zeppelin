package block

type CherryWood struct {
	Axis string
}

func (b CherryWood) Encode() (string, BlockProperties) {
	return "minecraft:cherry_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b CherryWood) New(props BlockProperties) Block {
	return CherryWood{
		Axis: props["axis"],
	}
}