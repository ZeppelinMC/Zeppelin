package block

type PolishedTuff struct {
}

func (b PolishedTuff) Encode() (string, BlockProperties) {
	return "minecraft:polished_tuff", BlockProperties{}
}

func (b PolishedTuff) New(props BlockProperties) Block {
	return PolishedTuff{}
}