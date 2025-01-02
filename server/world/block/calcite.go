package block

type Calcite struct {
}

func (b Calcite) Encode() (string, BlockProperties) {
	return "minecraft:calcite", BlockProperties{}
}

func (b Calcite) New(props BlockProperties) Block {
	return Calcite{}
}