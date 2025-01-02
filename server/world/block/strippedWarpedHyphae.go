package block

type StrippedWarpedHyphae struct {
	Axis string
}

func (b StrippedWarpedHyphae) Encode() (string, BlockProperties) {
	return "minecraft:stripped_warped_hyphae", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedWarpedHyphae) New(props BlockProperties) Block {
	return StrippedWarpedHyphae{
		Axis: props["axis"],
	}
}