package block

type EmeraldOre struct {
}

func (b EmeraldOre) Encode() (string, BlockProperties) {
	return "minecraft:emerald_ore", BlockProperties{}
}

func (b EmeraldOre) New(props BlockProperties) Block {
	return EmeraldOre{}
}