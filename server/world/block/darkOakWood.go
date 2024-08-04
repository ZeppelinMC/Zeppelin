package block

type DarkOakWood struct {
	Axis string
}

func (b DarkOakWood) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b DarkOakWood) New(props BlockProperties) Block {
	return DarkOakWood{
		Axis: props["axis"],
	}
}