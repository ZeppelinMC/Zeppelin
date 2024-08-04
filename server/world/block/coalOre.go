package block

type CoalOre struct {
}

func (b CoalOre) Encode() (string, BlockProperties) {
	return "minecraft:coal_ore", BlockProperties{}
}

func (b CoalOre) New(props BlockProperties) Block {
	return CoalOre{}
}