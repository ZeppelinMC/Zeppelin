package block

type BambooBlock struct {
	Axis string
}

func (b BambooBlock) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_block", BlockProperties{
		"axis": b.Axis,
	}
}

func (b BambooBlock) New(props BlockProperties) Block {
	return BambooBlock{
		Axis: props["axis"],
	}
}