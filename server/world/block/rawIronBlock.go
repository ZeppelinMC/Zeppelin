package block

type RawIronBlock struct {
}

func (b RawIronBlock) Encode() (string, BlockProperties) {
	return "minecraft:raw_iron_block", BlockProperties{}
}

func (b RawIronBlock) New(props BlockProperties) Block {
	return RawIronBlock{}
}