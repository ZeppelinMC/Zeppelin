package block

type Air struct {
}

func (b Air) Encode() (string, BlockProperties) {
	return "minecraft:air", BlockProperties{}
}

func (b Air) New(props BlockProperties) Block {
	return Air{}
}