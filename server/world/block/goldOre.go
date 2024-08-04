package block

type GoldOre struct {
}

func (b GoldOre) Encode() (string, BlockProperties) {
	return "minecraft:gold_ore", BlockProperties{}
}

func (b GoldOre) New(props BlockProperties) Block {
	return GoldOre{}
}