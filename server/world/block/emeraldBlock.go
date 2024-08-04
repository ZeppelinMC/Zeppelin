package block

type EmeraldBlock struct {
}

func (b EmeraldBlock) Encode() (string, BlockProperties) {
	return "minecraft:emerald_block", BlockProperties{}
}

func (b EmeraldBlock) New(props BlockProperties) Block {
	return EmeraldBlock{}
}