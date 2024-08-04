package block

type Sponge struct {
}

func (b Sponge) Encode() (string, BlockProperties) {
	return "minecraft:sponge", BlockProperties{}
}

func (b Sponge) New(props BlockProperties) Block {
	return Sponge{}
}