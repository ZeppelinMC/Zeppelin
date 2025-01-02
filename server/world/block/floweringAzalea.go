package block

type FloweringAzalea struct {
}

func (b FloweringAzalea) Encode() (string, BlockProperties) {
	return "minecraft:flowering_azalea", BlockProperties{}
}

func (b FloweringAzalea) New(props BlockProperties) Block {
	return FloweringAzalea{}
}