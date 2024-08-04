package block

type PolishedAndesite struct {
}

func (b PolishedAndesite) Encode() (string, BlockProperties) {
	return "minecraft:polished_andesite", BlockProperties{}
}

func (b PolishedAndesite) New(props BlockProperties) Block {
	return PolishedAndesite{}
}