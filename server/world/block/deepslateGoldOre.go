package block

type DeepslateGoldOre struct {
}

func (b DeepslateGoldOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_gold_ore", BlockProperties{}
}

func (b DeepslateGoldOre) New(props BlockProperties) Block {
	return DeepslateGoldOre{}
}