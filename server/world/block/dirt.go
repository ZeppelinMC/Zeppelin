package block

type Dirt struct {
}

func (b Dirt) Encode() (string, BlockProperties) {
	return "minecraft:dirt", BlockProperties{}
}

func (b Dirt) New(props BlockProperties) Block {
	return Dirt{}
}