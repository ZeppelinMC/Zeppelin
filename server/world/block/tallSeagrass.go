package block

type TallSeagrass struct {
	Half string
}

func (b TallSeagrass) Encode() (string, BlockProperties) {
	return "minecraft:tall_seagrass", BlockProperties{
		"half": b.Half,
	}
}

func (b TallSeagrass) New(props BlockProperties) Block {
	return TallSeagrass{
		Half: props["half"],
	}
}