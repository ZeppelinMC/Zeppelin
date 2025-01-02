package block

type PowderSnow struct {
}

func (b PowderSnow) Encode() (string, BlockProperties) {
	return "minecraft:powder_snow", BlockProperties{}
}

func (b PowderSnow) New(props BlockProperties) Block {
	return PowderSnow{}
}