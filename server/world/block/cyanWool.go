package block

type CyanWool struct {
}

func (b CyanWool) Encode() (string, BlockProperties) {
	return "minecraft:cyan_wool", BlockProperties{}
}

func (b CyanWool) New(props BlockProperties) Block {
	return CyanWool{}
}