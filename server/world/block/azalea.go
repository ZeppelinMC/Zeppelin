package block

type Azalea struct {
}

func (b Azalea) Encode() (string, BlockProperties) {
	return "minecraft:azalea", BlockProperties{}
}

func (b Azalea) New(props BlockProperties) Block {
	return Azalea{}
}