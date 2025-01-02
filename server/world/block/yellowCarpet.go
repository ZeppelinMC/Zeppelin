package block

type YellowCarpet struct {
}

func (b YellowCarpet) Encode() (string, BlockProperties) {
	return "minecraft:yellow_carpet", BlockProperties{}
}

func (b YellowCarpet) New(props BlockProperties) Block {
	return YellowCarpet{}
}