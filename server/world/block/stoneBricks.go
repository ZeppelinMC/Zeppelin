package block

type StoneBricks struct {
}

func (b StoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:stone_bricks", BlockProperties{}
}

func (b StoneBricks) New(props BlockProperties) Block {
	return StoneBricks{}
}