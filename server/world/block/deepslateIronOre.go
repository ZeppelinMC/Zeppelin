package block

type DeepslateIronOre struct {
}

func (b DeepslateIronOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_iron_ore", BlockProperties{}
}

func (b DeepslateIronOre) New(props BlockProperties) Block {
	return DeepslateIronOre{}
}