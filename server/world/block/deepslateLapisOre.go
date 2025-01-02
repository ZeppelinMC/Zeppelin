package block

type DeepslateLapisOre struct {
}

func (b DeepslateLapisOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_lapis_ore", BlockProperties{}
}

func (b DeepslateLapisOre) New(props BlockProperties) Block {
	return DeepslateLapisOre{}
}