package block

type LapisBlock struct {
}

func (b LapisBlock) Encode() (string, BlockProperties) {
	return "minecraft:lapis_block", BlockProperties{}
}

func (b LapisBlock) New(props BlockProperties) Block {
	return LapisBlock{}
}