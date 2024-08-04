package block

type CartographyTable struct {
}

func (b CartographyTable) Encode() (string, BlockProperties) {
	return "minecraft:cartography_table", BlockProperties{}
}

func (b CartographyTable) New(props BlockProperties) Block {
	return CartographyTable{}
}