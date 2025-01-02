package block

type RawCopperBlock struct {
}

func (b RawCopperBlock) Encode() (string, BlockProperties) {
	return "minecraft:raw_copper_block", BlockProperties{}
}

func (b RawCopperBlock) New(props BlockProperties) Block {
	return RawCopperBlock{}
}