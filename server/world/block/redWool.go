package block

type RedWool struct {
}

func (b RedWool) Encode() (string, BlockProperties) {
	return "minecraft:red_wool", BlockProperties{}
}

func (b RedWool) New(props BlockProperties) Block {
	return RedWool{}
}