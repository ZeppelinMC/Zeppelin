package block

type Clay struct {
}

func (b Clay) Encode() (string, BlockProperties) {
	return "minecraft:clay", BlockProperties{}
}

func (b Clay) New(props BlockProperties) Block {
	return Clay{}
}