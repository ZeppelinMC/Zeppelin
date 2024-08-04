package block

type ChiseledSandstone struct {
}

func (b ChiseledSandstone) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_sandstone", BlockProperties{}
}

func (b ChiseledSandstone) New(props BlockProperties) Block {
	return ChiseledSandstone{}
}