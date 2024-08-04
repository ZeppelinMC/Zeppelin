package block

type CrimsonHyphae struct {
	Axis string
}

func (b CrimsonHyphae) Encode() (string, BlockProperties) {
	return "minecraft:crimson_hyphae", BlockProperties{
		"axis": b.Axis,
	}
}

func (b CrimsonHyphae) New(props BlockProperties) Block {
	return CrimsonHyphae{
		Axis: props["axis"],
	}
}