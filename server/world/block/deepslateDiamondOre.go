package block

type DeepslateDiamondOre struct {
}

func (b DeepslateDiamondOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_diamond_ore", BlockProperties{}
}

func (b DeepslateDiamondOre) New(props BlockProperties) Block {
	return DeepslateDiamondOre{}
}