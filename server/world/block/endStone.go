package block

type EndStone struct {
}

func (b EndStone) Encode() (string, BlockProperties) {
	return "minecraft:end_stone", BlockProperties{}
}

func (b EndStone) New(props BlockProperties) Block {
	return EndStone{}
}