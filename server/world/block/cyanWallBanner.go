package block

type CyanWallBanner struct {
	Facing string
}

func (b CyanWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:cyan_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b CyanWallBanner) New(props BlockProperties) Block {
	return CyanWallBanner{
		Facing: props["facing"],
	}
}