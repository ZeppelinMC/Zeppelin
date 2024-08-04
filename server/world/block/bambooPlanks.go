package block

type BambooPlanks struct {
}

func (b BambooPlanks) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_planks", BlockProperties{}
}

func (b BambooPlanks) New(props BlockProperties) Block {
	return BambooPlanks{}
}