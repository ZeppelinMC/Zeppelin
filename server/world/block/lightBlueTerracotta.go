package block

type LightBlueTerracotta struct {
}

func (b LightBlueTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_terracotta", BlockProperties{}
}

func (b LightBlueTerracotta) New(props BlockProperties) Block {
	return LightBlueTerracotta{}
}