package block

type StrippedCrimsonStem struct {
	Axis string
}

func (b StrippedCrimsonStem) Encode() (string, BlockProperties) {
	return "minecraft:stripped_crimson_stem", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedCrimsonStem) New(props BlockProperties) Block {
	return StrippedCrimsonStem{
		Axis: props["axis"],
	}
}