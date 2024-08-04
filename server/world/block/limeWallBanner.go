package block

type LimeWallBanner struct {
	Facing string
}

func (b LimeWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:lime_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LimeWallBanner) New(props BlockProperties) Block {
	return LimeWallBanner{
		Facing: props["facing"],
	}
}