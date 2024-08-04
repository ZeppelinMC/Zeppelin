package block

type TuffBricks struct {
}

func (b TuffBricks) Encode() (string, BlockProperties) {
	return "minecraft:tuff_bricks", BlockProperties{}
}

func (b TuffBricks) New(props BlockProperties) Block {
	return TuffBricks{}
}