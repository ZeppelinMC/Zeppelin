package block

type DarkPrismarine struct {
}

func (b DarkPrismarine) Encode() (string, BlockProperties) {
	return "minecraft:dark_prismarine", BlockProperties{}
}

func (b DarkPrismarine) New(props BlockProperties) Block {
	return DarkPrismarine{}
}