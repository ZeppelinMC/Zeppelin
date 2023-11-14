package chunk

type Block interface {
	EncodedName() string

	New(string, map[string]string) Block

	Properties() map[string]string
}

type UnknownBlock struct {
	encodedName string
	properties  map[string]string
}

func (u UnknownBlock) EncodedName() string {
	return u.encodedName
}

func (u UnknownBlock) New(n string, m map[string]string) Block {
	return UnknownBlock{encodedName: n, properties: m}
}

func (u UnknownBlock) Properties() map[string]string {
	return u.properties
}

func NewUnknownBlock(name string) *UnknownBlock {
	return &UnknownBlock{encodedName: name}
}
