package block

type WitherRose struct {
}

func (b WitherRose) Encode() (string, BlockProperties) {
	return "minecraft:wither_rose", BlockProperties{}
}

func (b WitherRose) New(props BlockProperties) Block {
	return WitherRose{}
}