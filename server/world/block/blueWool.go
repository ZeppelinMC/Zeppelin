package block

type BlueWool struct {
}

func (b BlueWool) Encode() (string, BlockProperties) {
	return "minecraft:blue_wool", BlockProperties{}
}

func (b BlueWool) New(props BlockProperties) Block {
	return BlueWool{}
}