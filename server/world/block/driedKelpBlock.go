package block

type DriedKelpBlock struct {
}

func (b DriedKelpBlock) Encode() (string, BlockProperties) {
	return "minecraft:dried_kelp_block", BlockProperties{}
}

func (b DriedKelpBlock) New(props BlockProperties) Block {
	return DriedKelpBlock{}
}