package block

type OakWood struct {
	Axis string
}

func (b OakWood) Encode() (string, BlockProperties) {
	return "minecraft:oak_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b OakWood) New(props BlockProperties) Block {
	return OakWood{
		Axis: props["axis"],
	}
}