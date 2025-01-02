package block

type MangroveWood struct {
	Axis string
}

func (b MangroveWood) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b MangroveWood) New(props BlockProperties) Block {
	return MangroveWood{
		Axis: props["axis"],
	}
}