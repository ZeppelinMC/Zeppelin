package block

type LightBlueWool struct {
}

func (b LightBlueWool) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_wool", BlockProperties{}
}

func (b LightBlueWool) New(props BlockProperties) Block {
	return LightBlueWool{}
}