package block

type PinkWallBanner struct {
	Facing string
}

func (b PinkWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:pink_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PinkWallBanner) New(props BlockProperties) Block {
	return PinkWallBanner{
		Facing: props["facing"],
	}
}