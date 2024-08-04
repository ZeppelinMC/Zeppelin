package block

type AncientDebris struct {
}

func (b AncientDebris) Encode() (string, BlockProperties) {
	return "minecraft:ancient_debris", BlockProperties{}
}

func (b AncientDebris) New(props BlockProperties) Block {
	return AncientDebris{}
}