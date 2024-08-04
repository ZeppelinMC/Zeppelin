package block

type Obsidian struct {
}

func (b Obsidian) Encode() (string, BlockProperties) {
	return "minecraft:obsidian", BlockProperties{}
}

func (b Obsidian) New(props BlockProperties) Block {
	return Obsidian{}
}