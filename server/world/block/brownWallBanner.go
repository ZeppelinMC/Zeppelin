package block

type BrownWallBanner struct {
	Facing string
}

func (b BrownWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:brown_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BrownWallBanner) New(props BlockProperties) Block {
	return BrownWallBanner{
		Facing: props["facing"],
	}
}