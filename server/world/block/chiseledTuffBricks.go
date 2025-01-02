package block

type ChiseledTuffBricks struct {
}

func (b ChiseledTuffBricks) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_tuff_bricks", BlockProperties{}
}

func (b ChiseledTuffBricks) New(props BlockProperties) Block {
	return ChiseledTuffBricks{}
}