package block

type VerdantFroglight struct {
	Axis string
}

func (b VerdantFroglight) Encode() (string, BlockProperties) {
	return "minecraft:verdant_froglight", BlockProperties{
		"axis": b.Axis,
	}
}

func (b VerdantFroglight) New(props BlockProperties) Block {
	return VerdantFroglight{
		Axis: props["axis"],
	}
}