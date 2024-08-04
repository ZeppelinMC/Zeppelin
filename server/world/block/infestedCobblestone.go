package block

type InfestedCobblestone struct {
}

func (b InfestedCobblestone) Encode() (string, BlockProperties) {
	return "minecraft:infested_cobblestone", BlockProperties{}
}

func (b InfestedCobblestone) New(props BlockProperties) Block {
	return InfestedCobblestone{}
}