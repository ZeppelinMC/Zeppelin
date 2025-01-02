package block

type Bricks struct {
}

func (b Bricks) Encode() (string, BlockProperties) {
	return "minecraft:bricks", BlockProperties{}
}

func (b Bricks) New(props BlockProperties) Block {
	return Bricks{}
}