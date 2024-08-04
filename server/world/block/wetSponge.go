package block

type WetSponge struct {
}

func (b WetSponge) Encode() (string, BlockProperties) {
	return "minecraft:wet_sponge", BlockProperties{}
}

func (b WetSponge) New(props BlockProperties) Block {
	return WetSponge{}
}