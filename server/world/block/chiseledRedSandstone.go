package block

type ChiseledRedSandstone struct {
}

func (b ChiseledRedSandstone) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_red_sandstone", BlockProperties{}
}

func (b ChiseledRedSandstone) New(props BlockProperties) Block {
	return ChiseledRedSandstone{}
}