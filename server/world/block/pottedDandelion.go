package block

type PottedDandelion struct {
}

func (b PottedDandelion) Encode() (string, BlockProperties) {
	return "minecraft:potted_dandelion", BlockProperties{}
}

func (b PottedDandelion) New(props BlockProperties) Block {
	return PottedDandelion{}
}