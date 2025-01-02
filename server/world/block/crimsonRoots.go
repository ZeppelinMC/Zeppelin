package block

type CrimsonRoots struct {
}

func (b CrimsonRoots) Encode() (string, BlockProperties) {
	return "minecraft:crimson_roots", BlockProperties{}
}

func (b CrimsonRoots) New(props BlockProperties) Block {
	return CrimsonRoots{}
}