package block

type PottedCrimsonRoots struct {
}

func (b PottedCrimsonRoots) Encode() (string, BlockProperties) {
	return "minecraft:potted_crimson_roots", BlockProperties{}
}

func (b PottedCrimsonRoots) New(props BlockProperties) Block {
	return PottedCrimsonRoots{}
}