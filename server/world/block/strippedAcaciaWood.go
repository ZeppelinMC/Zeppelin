package block

type StrippedAcaciaWood struct {
	Axis string
}

func (b StrippedAcaciaWood) Encode() (string, BlockProperties) {
	return "minecraft:stripped_acacia_wood", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedAcaciaWood) New(props BlockProperties) Block {
	return StrippedAcaciaWood{
		Axis: props["axis"],
	}
}