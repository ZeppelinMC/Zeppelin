package block

type Dirt struct {
}

func (Dirt) Encode() (string, BlockProperties) {
	return "minecraft:dirt", nil
}

func (Dirt) New(props BlockProperties) Block {
	return Dirt{}
}

var _ Block = (*Dirt)(nil)
