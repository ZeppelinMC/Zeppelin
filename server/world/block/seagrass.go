package block

type Seagrass struct {
}

func (b Seagrass) Encode() (string, BlockProperties) {
	return "minecraft:seagrass", BlockProperties{}
}

func (b Seagrass) New(props BlockProperties) Block {
	return Seagrass{}
}