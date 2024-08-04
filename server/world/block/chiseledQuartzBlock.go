package block

type ChiseledQuartzBlock struct {
}

func (b ChiseledQuartzBlock) Encode() (string, BlockProperties) {
	return "minecraft:chiseled_quartz_block", BlockProperties{}
}

func (b ChiseledQuartzBlock) New(props BlockProperties) Block {
	return ChiseledQuartzBlock{}
}