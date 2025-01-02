package block

type MagentaWool struct {
}

func (b MagentaWool) Encode() (string, BlockProperties) {
	return "minecraft:magenta_wool", BlockProperties{}
}

func (b MagentaWool) New(props BlockProperties) Block {
	return MagentaWool{}
}