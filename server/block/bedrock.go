package block

type Bedrock struct{}

func (b Bedrock) EncodedName() string {
	return "minecraft:bedrock"
}

func (b Bedrock) New(m map[string]string) Block {
	return b
}

func (b Bedrock) Properties() map[string]string {
	return nil
}
