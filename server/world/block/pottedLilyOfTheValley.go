package block

type PottedLilyOfTheValley struct {
}

func (b PottedLilyOfTheValley) Encode() (string, BlockProperties) {
	return "minecraft:potted_lily_of_the_valley", BlockProperties{}
}

func (b PottedLilyOfTheValley) New(props BlockProperties) Block {
	return PottedLilyOfTheValley{}
}