package block

type ChiseledPolishedBlackstone struct {
}

func (b ChiseledPolishedBlackstone) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_polished_blackstone", BlockProperties{}
}

func (b ChiseledPolishedBlackstone) New(props BlockProperties) Block {
	return ChiseledPolishedBlackstone{}
}