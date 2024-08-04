package block

type WaxedCopperBlock struct {
}

func (b WaxedCopperBlock) Encode() (string, BlockProperties) {
	return "minecraft:waxed_copper_block", BlockProperties{}
}

func (b WaxedCopperBlock) New(props BlockProperties) Block {
	return WaxedCopperBlock{}
}