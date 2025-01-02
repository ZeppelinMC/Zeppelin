package block

type WarpedHyphae struct {
	Axis string
}

func (b WarpedHyphae) Encode() (string, BlockProperties) {
	return "minecraft:warped_hyphae", BlockProperties{
		"axis": b.Axis,
	}
}

func (b WarpedHyphae) New(props BlockProperties) Block {
	return WarpedHyphae{
		Axis: props["axis"],
	}
}