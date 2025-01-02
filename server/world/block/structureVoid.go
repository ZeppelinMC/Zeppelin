package block

type StructureVoid struct {
}

func (b StructureVoid) Encode() (string, BlockProperties) {
	return "minecraft:structure_void", BlockProperties{}
}

func (b StructureVoid) New(props BlockProperties) Block {
	return StructureVoid{}
}