package block

type GreenCarpet struct {
}

func (b GreenCarpet) Encode() (string, BlockProperties) {
	return "minecraft:green_carpet", BlockProperties{}
}

func (b GreenCarpet) New(props BlockProperties) Block {
	return GreenCarpet{}
}