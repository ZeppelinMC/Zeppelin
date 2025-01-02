package block

type SeaLantern struct {
}

func (b SeaLantern) Encode() (string, BlockProperties) {
	return "minecraft:sea_lantern", BlockProperties{}
}

func (b SeaLantern) New(props BlockProperties) Block {
	return SeaLantern{}
}