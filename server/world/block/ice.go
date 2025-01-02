package block

type Ice struct {
}

func (b Ice) Encode() (string, BlockProperties) {
	return "minecraft:ice", BlockProperties{}
}

func (b Ice) New(props BlockProperties) Block {
	return Ice{}
}