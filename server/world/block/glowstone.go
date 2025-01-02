package block

type Glowstone struct {
}

func (b Glowstone) Encode() (string, BlockProperties) {
	return "minecraft:glowstone", BlockProperties{}
}

func (b Glowstone) New(props BlockProperties) Block {
	return Glowstone{}
}