package block

type DiamondBlock struct {
}

func (b DiamondBlock) Encode() (string, BlockProperties) {
	return "minecraft:diamond_block", BlockProperties{}
}

func (b DiamondBlock) New(props BlockProperties) Block {
	return DiamondBlock{}
}