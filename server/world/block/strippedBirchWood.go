package block

type StrippedBirchWood struct {
	Axis string
}

func (b StrippedBirchWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_birch_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedBirchWood) New(props BlockProperties) Block {
	return StrippedBirchWood{
		Axis: props["axis"],
	}
}