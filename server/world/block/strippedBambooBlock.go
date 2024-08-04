package block

type StrippedBambooBlock struct {
	Axis string
}

func (b StrippedBambooBlock) Encode() (string, BlockProperties) {
	return "minecraft:stripped_bamboo_block", BlockProperties{
		"axis": b.Axis,
	}
}

func (b StrippedBambooBlock) New(props BlockProperties) Block {
	return StrippedBambooBlock{
		Axis: props["axis"],
	}
}