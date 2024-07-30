package block

type Dirt struct {
}

func (g Dirt) Encode() (string, BlockProperties) {
	return "minecraft:dirt", nil
}

func (g Dirt) New(props BlockProperties) Block {
	return Dirt{}
}

var _ Block = (*Dirt)(nil)
