package block

type LightBlueCarpet struct {
}

func (b LightBlueCarpet) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_carpet", BlockProperties{}
}

func (b LightBlueCarpet) New(props BlockProperties) Block {
	return LightBlueCarpet{}
}