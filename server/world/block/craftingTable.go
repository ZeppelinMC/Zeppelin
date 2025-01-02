package block

type CraftingTable struct {
}

func (b CraftingTable) Encode() (string, BlockProperties) {
	return "minecraft:crafting_table", BlockProperties{}
}

func (b CraftingTable) New(props BlockProperties) Block {
	return CraftingTable{}
}