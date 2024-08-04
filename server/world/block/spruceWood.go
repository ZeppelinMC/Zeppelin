package block

type SpruceWood struct {
	Axis string
}

func (b SpruceWood) Encode() (string, BlockProperties) {
	return "minecraft:spruce_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b SpruceWood) New(props BlockProperties) Block {
	return SpruceWood{
		Axis: props["axis"],
	}
}