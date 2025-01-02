package block

type LilyOfTheValley struct {
}

func (b LilyOfTheValley) Encode() (string, BlockProperties) {
	return "minecraft:lily_of_the_valley", BlockProperties{}
}

func (b LilyOfTheValley) New(props BlockProperties) Block {
	return LilyOfTheValley{}
}