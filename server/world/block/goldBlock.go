package block

type GoldBlock struct {
}

func (b GoldBlock) Encode() (string, BlockProperties) {
	return "minecraft:gold_block", BlockProperties{}
}

func (b GoldBlock) New(props BlockProperties) Block {
	return GoldBlock{}
}