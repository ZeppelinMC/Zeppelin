package block

type HayBlock struct {
	Axis string
}

func (b HayBlock) Encode() (string, BlockProperties) {
	return "minecraft:hay_block", BlockProperties{
		"axis": b.Axis,
	}
}

func (b HayBlock) New(props BlockProperties) Block {
	return HayBlock{
		Axis: props["axis"],
	}
}