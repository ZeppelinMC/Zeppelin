package block

type MossyCobblestone struct {
}

func (b MossyCobblestone) Encode() (string, BlockProperties) {
	return "minecraft:mossy_cobblestone", BlockProperties{}
}

func (b MossyCobblestone) New(props BlockProperties) Block {
	return MossyCobblestone{}
}