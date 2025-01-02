package block

type AcaciaWood struct {
	Axis string
}

func (b AcaciaWood) Encode() (string, BlockProperties) {
	return "minecraft:acacia_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b AcaciaWood) New(props BlockProperties) Block {
	return AcaciaWood{
		Axis: props["axis"],
	}
}