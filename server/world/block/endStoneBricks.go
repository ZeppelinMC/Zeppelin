package block

type EndStoneBricks struct {
}

func (b EndStoneBricks) Encode() (string, BlockProperties) {
	return "minecraft:end_stone_bricks", BlockProperties{}
}

func (b EndStoneBricks) New(props BlockProperties) Block {
	return EndStoneBricks{}
}