package block

type IronOre struct {
}

func (b IronOre) Encode() (string, BlockProperties) {
	return "minecraft:iron_ore", BlockProperties{}
}

func (b IronOre) New(props BlockProperties) Block {
	return IronOre{}
}