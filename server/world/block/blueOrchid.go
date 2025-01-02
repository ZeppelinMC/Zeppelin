package block

type BlueOrchid struct {
}

func (b BlueOrchid) Encode() (string, BlockProperties) {
	return "minecraft:blue_orchid", BlockProperties{}
}

func (b BlueOrchid) New(props BlockProperties) Block {
	return BlueOrchid{}
}