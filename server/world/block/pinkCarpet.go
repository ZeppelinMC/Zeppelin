package block

type PinkCarpet struct {
}

func (b PinkCarpet) Encode() (string, BlockProperties) {
	return "minecraft:pink_carpet", BlockProperties{}
}

func (b PinkCarpet) New(props BlockProperties) Block {
	return PinkCarpet{}
}