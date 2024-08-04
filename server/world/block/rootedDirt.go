package block

type RootedDirt struct {
}

func (b RootedDirt) Encode() (string, BlockProperties) {
	return "minecraft:rooted_dirt", BlockProperties{}
}

func (b RootedDirt) New(props BlockProperties) Block {
	return RootedDirt{}
}