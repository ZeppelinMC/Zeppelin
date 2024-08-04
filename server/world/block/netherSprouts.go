package block

type NetherSprouts struct {
}

func (b NetherSprouts) Encode() (string, BlockProperties) {
	return "minecraft:nether_sprouts", BlockProperties{}
}

func (b NetherSprouts) New(props BlockProperties) Block {
	return NetherSprouts{}
}