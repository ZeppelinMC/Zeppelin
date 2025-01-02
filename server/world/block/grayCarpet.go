package block

type GrayCarpet struct {
}

func (b GrayCarpet) Encode() (string, BlockProperties) {
	return "minecraft:gray_carpet", BlockProperties{}
}

func (b GrayCarpet) New(props BlockProperties) Block {
	return GrayCarpet{}
}