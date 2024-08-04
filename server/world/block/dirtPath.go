package block

type DirtPath struct {
}

func (b DirtPath) Encode() (string, BlockProperties) {
	return "minecraft:dirt_path", BlockProperties{}
}

func (b DirtPath) New(props BlockProperties) Block {
	return DirtPath{}
}