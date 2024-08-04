package block

type QuartzBricks struct {
}

func (b QuartzBricks) Encode() (string, BlockProperties) {
	return "minecraft:quartz_bricks", BlockProperties{}
}

func (b QuartzBricks) New(props BlockProperties) Block {
	return QuartzBricks{}
}