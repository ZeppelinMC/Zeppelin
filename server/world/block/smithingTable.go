package block

type SmithingTable struct {
}

func (b SmithingTable) Encode() (string, BlockProperties) {
	return "minecraft:smithing_table", BlockProperties{}
}

func (b SmithingTable) New(props BlockProperties) Block {
	return SmithingTable{}
}