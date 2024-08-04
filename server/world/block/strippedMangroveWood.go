package block

type StrippedMangroveWood struct {
	Axis string
}

func (b StrippedMangroveWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_mangrove_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedMangroveWood) New(props BlockProperties) Block {
	return StrippedMangroveWood{
		Axis: props["axis"],
	}
}