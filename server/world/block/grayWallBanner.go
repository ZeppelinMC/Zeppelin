package block

type GrayWallBanner struct {
	Facing string
}

func (b GrayWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:gray_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GrayWallBanner) New(props BlockProperties) Block {
	return GrayWallBanner{
		Facing: props["facing"],
	}
}