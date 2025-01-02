package block

type DeepslateCoalOre struct {
}

func (b DeepslateCoalOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_coal_ore", BlockProperties{}
}

func (b DeepslateCoalOre) New(props BlockProperties) Block {
	return DeepslateCoalOre{}
}