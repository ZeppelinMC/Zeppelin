package block

type WarpedRoots struct {
}

func (b WarpedRoots) Encode() (string, BlockProperties) {
	return "minecraft:warped_roots", BlockProperties{}
}

func (b WarpedRoots) New(props BlockProperties) Block {
	return WarpedRoots{}
}