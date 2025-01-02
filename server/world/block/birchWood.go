package block

type BirchWood struct {
	Axis string
}

func (b BirchWood) Encode() (string, BlockProperties) {
	return "minecraft:birch_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b BirchWood) New(props BlockProperties) Block {
	return BirchWood{
		Axis: props["axis"],
	}
}