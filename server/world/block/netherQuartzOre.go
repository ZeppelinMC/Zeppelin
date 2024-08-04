package block

type NetherQuartzOre struct {
}

func (b NetherQuartzOre) Encode() (string, BlockProperties) {
	return "minecraft:nether_quartz_ore", BlockProperties{}
}

func (b NetherQuartzOre) New(props BlockProperties) Block {
	return NetherQuartzOre{}
}