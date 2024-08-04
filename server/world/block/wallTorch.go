package block

type WallTorch struct {
	Facing string
}

func (b WallTorch) Encode() (string, BlockProperties) {
	return "minecraft:wall_torch", BlockProperties{
		"facing": b.Facing,
	}
}

func (b WallTorch) New(props BlockProperties) Block {
	return WallTorch{
		Facing: props["facing"],
	}
}