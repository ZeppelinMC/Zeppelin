package block

type SlimeBlock struct {
}

func (b SlimeBlock) Encode() (string, BlockProperties) {
	return "minecraft:slime_block", BlockProperties{}
}

func (b SlimeBlock) New(props BlockProperties) Block {
	return SlimeBlock{}
}