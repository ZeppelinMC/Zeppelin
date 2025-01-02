package block

type OrangeWool struct {
}

func (b OrangeWool) Encode() (string, BlockProperties) {
	return "minecraft:orange_wool", BlockProperties{}
}

func (b OrangeWool) New(props BlockProperties) Block {
	return OrangeWool{}
}