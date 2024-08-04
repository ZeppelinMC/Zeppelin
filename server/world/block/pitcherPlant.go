package block

type PitcherPlant struct {
	Half string
}

func (b PitcherPlant) Encode() (string, BlockProperties) {
	return "minecraft:pitcher_plant", BlockProperties{
		"half": b.Half,
	}
}

func (b PitcherPlant) New(props BlockProperties) Block {
	return PitcherPlant{
		Half: props["half"],
	}
}