package block

type BlackWallBanner struct {
	Facing string
}

func (b BlackWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:black_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlackWallBanner) New(props BlockProperties) Block {
	return BlackWallBanner{
		Facing: props["facing"],
	}
}