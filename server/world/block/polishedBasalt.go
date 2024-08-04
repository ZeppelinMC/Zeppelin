package block

type PolishedBasalt struct {
	Axis string
}

func (b PolishedBasalt) Encode() (string, BlockProperties) {
	return "minecraft:polished_basalt", BlockProperties{
		"axis": b.Axis,
	}
}

func (b PolishedBasalt) New(props BlockProperties) Block {
	return PolishedBasalt{
		Axis: props["axis"],
	}
}