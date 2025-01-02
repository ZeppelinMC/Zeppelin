package block

type BrainCoralBlock struct {
}

func (b BrainCoralBlock) Encode() (string, BlockProperties) {
	return "minecraft:brain_coral_block", BlockProperties{}
}

func (b BrainCoralBlock) New(props BlockProperties) Block {
	return BrainCoralBlock{}
}