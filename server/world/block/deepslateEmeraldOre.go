package block

type DeepslateEmeraldOre struct {
}

func (b DeepslateEmeraldOre) Encode() (string, BlockProperties) {
	return "minecraft:deepslate_emerald_ore", BlockProperties{}
}

func (b DeepslateEmeraldOre) New(props BlockProperties) Block {
	return DeepslateEmeraldOre{}
}