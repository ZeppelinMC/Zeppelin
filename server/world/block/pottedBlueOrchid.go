package block

type PottedBlueOrchid struct {
}

func (b PottedBlueOrchid) Encode() (string, BlockProperties) {
	return "minecraft:potted_blue_orchid", BlockProperties{}
}

func (b PottedBlueOrchid) New(props BlockProperties) Block {
	return PottedBlueOrchid{}
}