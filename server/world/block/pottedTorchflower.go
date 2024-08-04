package block

type PottedTorchflower struct {
}

func (b PottedTorchflower) Encode() (string, BlockProperties) {
	return "minecraft:potted_torchflower", BlockProperties{}
}

func (b PottedTorchflower) New(props BlockProperties) Block {
	return PottedTorchflower{}
}