package block

type SoulWallTorch struct {
	Facing string
}

func (b SoulWallTorch) Encode() (string, BlockProperties) {
	return "minecraft:soul_wall_torch", BlockProperties{
		"facing": b.Facing,
	}
}

func (b SoulWallTorch) New(props BlockProperties) Block {
	return SoulWallTorch{
		Facing: props["facing"],
	}
}