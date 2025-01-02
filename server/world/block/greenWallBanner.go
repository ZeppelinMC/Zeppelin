package block

type GreenWallBanner struct {
	Facing string
}

func (b GreenWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:green_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GreenWallBanner) New(props BlockProperties) Block {
	return GreenWallBanner{
		Facing: props["facing"],
	}
}