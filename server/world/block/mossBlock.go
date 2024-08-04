package block

type MossBlock struct {
}

func (b MossBlock) Encode() (string, BlockProperties) {
	return "minecraft:moss_block", BlockProperties{}
}

func (b MossBlock) New(props BlockProperties) Block {
	return MossBlock{}
}