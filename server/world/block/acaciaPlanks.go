package block

type AcaciaPlanks struct {
}

func (b AcaciaPlanks) Encode() (string, BlockProperties) {
	return "minecraft:acacia_planks", BlockProperties{}
}

func (b AcaciaPlanks) New(props BlockProperties) Block {
	return AcaciaPlanks{}
}