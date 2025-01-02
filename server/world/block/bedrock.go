package block

type Bedrock struct {
}

func (b Bedrock) Encode() (string, BlockProperties) {
	return "minecraft:bedrock", BlockProperties{}
}

func (b Bedrock) New(props BlockProperties) Block {
	return Bedrock{}
}