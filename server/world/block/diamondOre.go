package block

type DiamondOre struct {
}

func (b DiamondOre) Encode() (string, BlockProperties) {
	return "minecraft:diamond_ore", BlockProperties{}
}

func (b DiamondOre) New(props BlockProperties) Block {
	return DiamondOre{}
}