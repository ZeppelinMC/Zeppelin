package block

type JungleWood struct {
	Axis string
}

func (b JungleWood) Encode() (string, BlockProperties) {
	return "minecraft:jungle_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b JungleWood) New(props BlockProperties) Block {
	return JungleWood{
		Axis: props["axis"],
	}
}