package block

type AmethystBlock struct {
}

func (b AmethystBlock) Encode() (string, BlockProperties) {
	return "minecraft:amethyst_block", BlockProperties{}
}

func (b AmethystBlock) New(props BlockProperties) Block {
	return AmethystBlock{}
}