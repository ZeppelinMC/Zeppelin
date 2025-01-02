package block

type BlueCarpet struct {
}

func (b BlueCarpet) Encode() (string, BlockProperties) {
	return "minecraft:blue_carpet", BlockProperties{}
}

func (b BlueCarpet) New(props BlockProperties) Block {
	return BlueCarpet{}
}