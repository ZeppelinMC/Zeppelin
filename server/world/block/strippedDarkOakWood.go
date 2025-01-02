package block

type StrippedDarkOakWood struct {
	Axis string
}

func (b StrippedDarkOakWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_dark_oak_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedDarkOakWood) New(props BlockProperties) Block {
	return StrippedDarkOakWood{
		Axis: props["axis"],
	}
}