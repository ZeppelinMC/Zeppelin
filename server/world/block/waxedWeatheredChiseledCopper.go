package block

type WaxedWeatheredChiseledCopper struct {
}

func (b WaxedWeatheredChiseledCopper) Encode() (string, BlockProperties) {
	return "minecraft:waxed_weathered_chiseled_copper", BlockProperties{}
}

func (b WaxedWeatheredChiseledCopper) New(props BlockProperties) Block {
	return WaxedWeatheredChiseledCopper{}
}