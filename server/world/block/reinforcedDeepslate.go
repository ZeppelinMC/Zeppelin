package block

type ReinforcedDeepslate struct {
}

func (b ReinforcedDeepslate) Encode() (string, BlockProperties) {
	return "minecraft:reinforced_deepslate", BlockProperties{}
}

func (b ReinforcedDeepslate) New(props BlockProperties) Block {
	return ReinforcedDeepslate{}
}