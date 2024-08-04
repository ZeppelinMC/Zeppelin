package block

type CyanCarpet struct {
}

func (b CyanCarpet) Encode() (string, BlockProperties) {
	return "minecraft:cyan_carpet", BlockProperties{}
}

func (b CyanCarpet) New(props BlockProperties) Block {
	return CyanCarpet{}
}