package block

type PolishedBlackstone struct {
}

func (b PolishedBlackstone) Encode() (string, BlockProperties) {
	return "minecraft:polished_blackstone", BlockProperties{}
}

func (b PolishedBlackstone) New(props BlockProperties) Block {
	return PolishedBlackstone{}
}