package block

type CopperOre struct {
}

func (b CopperOre) Encode() (string, BlockProperties) {
	return "minecraft:copper_ore", BlockProperties{}
}

func (b CopperOre) New(props BlockProperties) Block {
	return CopperOre{}
}