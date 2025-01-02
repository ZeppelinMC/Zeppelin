package block

type BrownCarpet struct {
}

func (b BrownCarpet) Encode() (string, BlockProperties) {
	return "minecraft:brown_carpet", BlockProperties{}
}

func (b BrownCarpet) New(props BlockProperties) Block {
	return BrownCarpet{}
}