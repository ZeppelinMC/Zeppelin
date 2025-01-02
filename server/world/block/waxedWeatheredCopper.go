package block

type WaxedWeatheredCopper struct {
}

func (b WaxedWeatheredCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_copper", BlockProperties{}
}

func (b WaxedWeatheredCopper) New(props BlockProperties) Block {
	return WaxedWeatheredCopper{}
}