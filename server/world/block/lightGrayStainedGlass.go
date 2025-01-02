package block

type LightGrayStainedGlass struct {
}

func (b LightGrayStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_stained_glass", BlockProperties{}
}

func (b LightGrayStainedGlass) New(props BlockProperties) Block {
	return LightGrayStainedGlass{}
}