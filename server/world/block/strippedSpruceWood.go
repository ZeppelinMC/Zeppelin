package block

type StrippedSpruceWood struct {
	Axis string
}

func (b StrippedSpruceWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_spruce_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedSpruceWood) New(props BlockProperties) Block {
	return StrippedSpruceWood{
		Axis: props["axis"],
	}
}