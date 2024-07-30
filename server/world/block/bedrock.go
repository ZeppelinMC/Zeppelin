package block

type Bedrock struct {
}

func (g Bedrock) Encode() (string, BlockProperties) {
	return "minecraft:bedrock", nil
}

func (g Bedrock) New(props BlockProperties) Block {
	return Dirt{}
}

var _ Block = (*Bedrock)(nil)
