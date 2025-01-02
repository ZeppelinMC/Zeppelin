package block

type CoarseDirt struct {
}

func (b CoarseDirt) Encode() (string, BlockProperties) {
	return "minecraft:coarse_dirt", BlockProperties{}
}

func (b CoarseDirt) New(props BlockProperties) Block {
	return CoarseDirt{}
}