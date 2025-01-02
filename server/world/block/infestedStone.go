package block

type InfestedStone struct {
}

func (b InfestedStone) Encode() (string, BlockProperties) {
	return "minecraft:infested_stone", BlockProperties{}
}

func (b InfestedStone) New(props BlockProperties) Block {
	return InfestedStone{}
}