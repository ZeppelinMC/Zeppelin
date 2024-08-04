package block

type CopperBlock struct {
}

func (b CopperBlock) Encode() (string, BlockProperties) {
	return "minecraft:copper_block", BlockProperties{}
}

func (b CopperBlock) New(props BlockProperties) Block {
	return CopperBlock{}
}