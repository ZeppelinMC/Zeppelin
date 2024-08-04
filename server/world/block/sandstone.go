package block

type Sandstone struct {
}

func (b Sandstone) Encode() (string, BlockProperties) {
	return "minecraft:sandstone", BlockProperties{}
}

func (b Sandstone) New(props BlockProperties) Block {
	return Sandstone{}
}