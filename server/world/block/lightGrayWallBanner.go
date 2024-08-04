package block

type LightGrayWallBanner struct {
	Facing string
}

func (b LightGrayWallBanner) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_wall_banner", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightGrayWallBanner) New(props BlockProperties) Block {
	return LightGrayWallBanner{
		Facing: props["facing"],
	}
}