package block

type FletchingTable struct {
}

func (b FletchingTable) Encode() (string, BlockProperties) {
	return "minecraft:fletching_table", BlockProperties{}
}

func (b FletchingTable) New(props BlockProperties) Block {
	return FletchingTable{}
}