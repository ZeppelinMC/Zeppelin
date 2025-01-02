package block

type RawGoldBlock struct {
}

func (b RawGoldBlock) Encode() (string, BlockProperties) {
	return "minecraft:raw_gold_block", BlockProperties{}
}

func (b RawGoldBlock) New(props BlockProperties) Block {
	return RawGoldBlock{}
}