package block

type WhiteCarpet struct {
}

func (b WhiteCarpet) Encode() (string, BlockProperties) {
	return "minecraft:white_carpet", BlockProperties{}
}

func (b WhiteCarpet) New(props BlockProperties) Block {
	return WhiteCarpet{}
}