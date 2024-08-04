package block

type StrippedCherryWood struct {
	Axis string
}

func (b StrippedCherryWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_cherry_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedCherryWood) New(props BlockProperties) Block {
	return StrippedCherryWood{
		Axis: props["axis"],
	}
}