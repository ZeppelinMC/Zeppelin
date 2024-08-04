package block

type LimeWool struct {
}

func (b LimeWool) Encode() (string, BlockProperties) {
	return "minecraft:lime_wool", BlockProperties{}
}

func (b LimeWool) New(props BlockProperties) Block {
	return LimeWool{}
}