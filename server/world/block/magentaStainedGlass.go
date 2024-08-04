package block

type MagentaStainedGlass struct {
}

func (b MagentaStainedGlass) Encode() (string, BlockProperties) {
	return "minecraft:magenta_stained_glass", BlockProperties{}
}

func (b MagentaStainedGlass) New(props BlockProperties) Block {
	return MagentaStainedGlass{}
}