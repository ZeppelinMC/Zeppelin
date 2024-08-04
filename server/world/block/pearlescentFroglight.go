package block

type PearlescentFroglight struct {
	Axis string
}

func (b PearlescentFroglight) Encode() (string, BlockProperties) {
	return "minecraft:pearlescent_froglight", BlockProperties{
		"axis": b.Axis,
	}
}

func (b PearlescentFroglight) New(props BlockProperties) Block {
	return PearlescentFroglight{
		Axis: props["axis"],
	}
}