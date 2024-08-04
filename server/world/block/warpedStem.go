package block

type WarpedStem struct {
	Axis string
}

func (b WarpedStem) Encode() (string, BlockProperties) {
	return "minecraft:warped_stem", BlockProperties{
		"axis": b.Axis,
	}
}

func (b WarpedStem) New(props BlockProperties) Block {
	return WarpedStem{
		Axis: props["axis"],
	}
}