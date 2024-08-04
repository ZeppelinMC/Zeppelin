package block

type Spawner struct {
}

func (b Spawner) Encode() (string, BlockProperties) {
	return "minecraft:spawner", BlockProperties{}
}

func (b Spawner) New(props BlockProperties) Block {
	return Spawner{}
}