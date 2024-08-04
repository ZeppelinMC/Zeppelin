package block

type SporeBlossom struct {
}

func (b SporeBlossom) Encode() (string, BlockProperties) {
	return "minecraft:spore_blossom", BlockProperties{}
}

func (b SporeBlossom) New(props BlockProperties) Block {
	return SporeBlossom{}
}