package block

type MagentaTerracotta struct {
}

func (b MagentaTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:magenta_terracotta", BlockProperties{}
}

func (b MagentaTerracotta) New(props BlockProperties) Block {
	return MagentaTerracotta{}
}