package block

type NetherGoldOre struct {
}

func (b NetherGoldOre) Encode() (string, BlockProperties) {
	return "minecraft:nether_gold_ore", BlockProperties{}
}

func (b NetherGoldOre) New(props BlockProperties) Block {
	return NetherGoldOre{}
}