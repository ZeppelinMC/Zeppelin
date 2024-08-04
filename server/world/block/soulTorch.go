package block

type SoulTorch struct {
}

func (b SoulTorch) Encode() (string, BlockProperties) {
	return "minecraft:soul_torch", BlockProperties{}
}

func (b SoulTorch) New(props BlockProperties) Block {
	return SoulTorch{}
}