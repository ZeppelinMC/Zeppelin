package block

type LightGrayCarpet struct {
}

func (b LightGrayCarpet) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_carpet", BlockProperties{}
}

func (b LightGrayCarpet) New(props BlockProperties) Block {
	return LightGrayCarpet{}
}