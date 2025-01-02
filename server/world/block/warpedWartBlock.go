package block

type WarpedWartBlock struct {
}

func (b WarpedWartBlock) Encode() (string, BlockProperties) {
	return "minecraft:warped_wart_block", BlockProperties{}
}

func (b WarpedWartBlock) New(props BlockProperties) Block {
	return WarpedWartBlock{}
}