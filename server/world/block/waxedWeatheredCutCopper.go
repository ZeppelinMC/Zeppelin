package block

type WaxedWeatheredCutCopper struct {
}

func (b WaxedWeatheredCutCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_cut_copper", BlockProperties{}
}

func (b WaxedWeatheredCutCopper) New(props BlockProperties) Block {
	return WaxedWeatheredCutCopper{}
}