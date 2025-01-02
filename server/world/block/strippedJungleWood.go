package block

type StrippedJungleWood struct {
	Axis string
}

func (b StrippedJungleWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_jungle_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedJungleWood) New(props BlockProperties) Block {
	return StrippedJungleWood{
		Axis: props["axis"],
	}
}