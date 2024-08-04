package block

type PurpleStainedGlass struct {
}

func (b PurpleStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:purple_stained_glass", BlockProperties{}
}

func (b PurpleStainedGlass) New(props BlockProperties) Block {
	return PurpleStainedGlass{}
}