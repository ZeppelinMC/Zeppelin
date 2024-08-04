package block

type NetheriteBlock struct {
}

func (b NetheriteBlock) Encode() (string, BlockProperties) {
	return "minecraft:netherite_block", BlockProperties{}
}

func (b NetheriteBlock) New(props BlockProperties) Block {
	return NetheriteBlock{}
}