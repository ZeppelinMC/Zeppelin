package block

type Bedrock struct {
}

func (Bedrock) Encode() (string, BlockProperties) {
	return "minecraft:bedrock", nil
}

func (Bedrock) New(props BlockProperties) Block {
	return Bedrock{}
}

var _ Block = (*Bedrock)(nil)
