package block

type TintedGlass struct {
}

func (b TintedGlass) Encode() (string, BlockProperties) {
	return "minecraft:tinted_glass", BlockProperties{}
}

func (b TintedGlass) New(props BlockProperties) Block {
	return TintedGlass{}
}