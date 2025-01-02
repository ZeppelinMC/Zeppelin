package block

type WeatheredCopper struct {
}

func (b WeatheredCopper) Encode() (string, BlockProperties) {
	return "minecraft:weathered_copper", BlockProperties{}
}

func (b WeatheredCopper) New(props BlockProperties) Block {
	return WeatheredCopper{}
}