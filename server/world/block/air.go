package block

type Air struct {
}

func (Air) Encode() (string, BlockProperties) {
	return "minecraft:air", nil
}

func (Air) New(BlockProperties) Block {
	return Air{}
}

var _ Block = (*Air)(nil)
