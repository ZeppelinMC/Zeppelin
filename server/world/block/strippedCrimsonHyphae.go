package block

type StrippedCrimsonHyphae struct {
	Axis string
}

func (b StrippedCrimsonHyphae) Encode() (string, BlockProperties) {
	return "minecraft:stripped_crimson_hyphae", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedCrimsonHyphae) New(props BlockProperties) Block {
	return StrippedCrimsonHyphae{
		Axis: props["axis"],
	}
}