package block

type LightGrayTerracotta struct {
}

func (b LightGrayTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_terracotta", BlockProperties{}
}

func (b LightGrayTerracotta) New(props BlockProperties) Block {
	return LightGrayTerracotta{}
}