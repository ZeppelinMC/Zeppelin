package block

type PolishedDeepslate struct {
}

func (b PolishedDeepslate) Encode() (string, BlockProperties) {
	return "minecraft:polished_deepslate", BlockProperties{}
}

func (b PolishedDeepslate) New(props BlockProperties) Block {
	return PolishedDeepslate{}
}