package block

type Shroomlight struct {
}

func (b Shroomlight) Encode() (string, BlockProperties) {
	return "minecraft:shroomlight", BlockProperties{}
}

func (b Shroomlight) New(props BlockProperties) Block {
	return Shroomlight{}
}