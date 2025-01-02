package block

type PottedWarpedRoots struct {
}

func (b PottedWarpedRoots) Encode() (string, BlockProperties) {
	return "minecraft:potted_warped_roots", BlockProperties{}
}

func (b PottedWarpedRoots) New(props BlockProperties) Block {
	return PottedWarpedRoots{}
}