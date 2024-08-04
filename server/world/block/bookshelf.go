package block

type Bookshelf struct {
}

func (b Bookshelf) Encode() (string, BlockProperties) {
	return "minecraft:bookshelf", BlockProperties{}
}

func (b Bookshelf) New(props BlockProperties) Block {
	return Bookshelf{}
}