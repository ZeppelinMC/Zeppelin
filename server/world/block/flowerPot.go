package block

type FlowerPot struct {
}

func (b FlowerPot) Encode() (string, BlockProperties) {
	return "minecraft:flower_pot", BlockProperties{}
}

func (b FlowerPot) New(props BlockProperties) Block {
	return FlowerPot{}
}