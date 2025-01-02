package block

type PurpleCarpet struct {
}

func (b PurpleCarpet) Encode() (string, BlockProperties) {
	return "minecraft:purple_carpet", BlockProperties{}
}

func (b PurpleCarpet) New(props BlockProperties) Block {
	return PurpleCarpet{}
}