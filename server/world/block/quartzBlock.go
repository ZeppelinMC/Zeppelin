package block

type QuartzBlock struct {
}

func (b QuartzBlock) Encode() (string, BlockProperties) {
	return "minecraft:quartz_block", BlockProperties{}
}

func (b QuartzBlock) New(props BlockProperties) Block {
	return QuartzBlock{}
}