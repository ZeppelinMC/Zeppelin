package block

type StrippedWarpedStem struct {
	Axis string
}

func (b StrippedWarpedStem) Encode() (string, BlockProperties) {
	return "minecraft:stripped_warped_stem", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedWarpedStem) New(props BlockProperties) Block {
	return StrippedWarpedStem{
		Axis: props["axis"],
	}
}