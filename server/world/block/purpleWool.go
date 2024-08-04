package block

type PurpleWool struct {
}

func (b PurpleWool) Encode() (string, BlockProperties) {
	return "minecraft:purple_wool", BlockProperties{}
}

func (b PurpleWool) New(props BlockProperties) Block {
	return PurpleWool{}
}