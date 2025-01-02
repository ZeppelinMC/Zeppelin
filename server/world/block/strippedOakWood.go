package block

type StrippedOakWood struct {
	Axis string
}

func (b StrippedOakWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_oak_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedOakWood) New(props BlockProperties) Block {
	return StrippedOakWood{
		Axis: props["axis"],
	}
}