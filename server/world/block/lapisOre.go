package block

type LapisOre struct {
}

func (b LapisOre) Encode() (string, BlockProperties) {
	return "minecraft:lapis_ore", BlockProperties{}
}

func (b LapisOre) New(props BlockProperties) Block {
	return LapisOre{}
}