package block

type Frogspawn struct {
}

func (b Frogspawn) Encode() (string, BlockProperties) {
	return "minecraft:frogspawn", BlockProperties{}
}

func (b Frogspawn) New(props BlockProperties) Block {
	return Frogspawn{}
}