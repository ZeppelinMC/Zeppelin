package block

type HoneyBlock struct {
}

func (b HoneyBlock) Encode() (string, BlockProperties) {
	return "minecraft:honey_block", BlockProperties{}
}

func (b HoneyBlock) New(props BlockProperties) Block {
	return HoneyBlock{}
}