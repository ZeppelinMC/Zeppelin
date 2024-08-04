package block

type DeepslateCopperOre struct {
}

func (b DeepslateCopperOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_copper_ore", BlockProperties{}
}

func (b DeepslateCopperOre) New(props BlockProperties) Block {
	return DeepslateCopperOre{}
}