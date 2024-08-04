package block

type CobbledDeepslate struct {
}

func (b CobbledDeepslate) Encode() (string, BlockProperties) {
	return "minecraft:cobbled_deepslate", BlockProperties{}
}

func (b CobbledDeepslate) New(props BlockProperties) Block {
	return CobbledDeepslate{}
}