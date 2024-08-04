package block

type CrimsonNylium struct {
}

func (b CrimsonNylium) Encode() (string, BlockProperties) {
	return "minecraft:crimson_nylium", BlockProperties{}
}

func (b CrimsonNylium) New(props BlockProperties) Block {
	return CrimsonNylium{}
}