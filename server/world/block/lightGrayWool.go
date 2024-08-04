package block

type LightGrayWool struct {
}

func (b LightGrayWool) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_wool", BlockProperties{}
}

func (b LightGrayWool) New(props BlockProperties) Block {
	return LightGrayWool{}
}