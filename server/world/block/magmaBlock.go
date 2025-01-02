package block

type MagmaBlock struct {
}

func (b MagmaBlock) Encode() (string, BlockProperties) {
	return "minecraft:magma_block", BlockProperties{}
}

func (b MagmaBlock) New(props BlockProperties) Block {
	return MagmaBlock{}
}