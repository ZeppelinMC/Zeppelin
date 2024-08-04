package block

type ChiseledDeepslate struct {
}

func (b ChiseledDeepslate) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_deepslate", BlockProperties{}
}

func (b ChiseledDeepslate) New(props BlockProperties) Block {
	return ChiseledDeepslate{}
}