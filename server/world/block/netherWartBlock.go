package block

type NetherWartBlock struct {
}

func (b NetherWartBlock) Encode() (string, BlockProperties) {
	return "minecraft:nether_wart_block", BlockProperties{}
}

func (b NetherWartBlock) New(props BlockProperties) Block {
	return NetherWartBlock{}
}