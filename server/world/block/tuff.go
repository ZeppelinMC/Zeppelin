package block

type Tuff struct {
}

func (b Tuff) Encode() (string, BlockProperties) {
	return "minecraft:tuff", BlockProperties{}
}

func (b Tuff) New(props BlockProperties) Block {
	return Tuff{}
}