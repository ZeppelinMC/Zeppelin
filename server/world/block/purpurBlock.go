package block

type PurpurBlock struct {
}

func (b PurpurBlock) Encode() (string, BlockProperties) {
	return "minecraft:purpur_block", BlockProperties{}
}

func (b PurpurBlock) New(props BlockProperties) Block {
	return PurpurBlock{}
}