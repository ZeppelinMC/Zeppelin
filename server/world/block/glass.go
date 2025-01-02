package block

type Glass struct {
}

func (b Glass) Encode() (string, BlockProperties) {
	return "minecraft:glass", BlockProperties{}
}

func (b Glass) New(props BlockProperties) Block {
	return Glass{}
}