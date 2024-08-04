package block

type CrimsonFungus struct {
}

func (b CrimsonFungus) Encode() (string, BlockProperties) {
	return "minecraft:crimson_fungus", BlockProperties{}
}

func (b CrimsonFungus) New(props BlockProperties) Block {
	return CrimsonFungus{}
}