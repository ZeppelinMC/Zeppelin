package block

type UnknownBlock struct {
	encodedName string
	properties  map[string]string
}

func (u UnknownBlock) EncodedName() string {
	return u.encodedName
}

func (u UnknownBlock) New(m map[string]string) Block {
	return UnknownBlock{encodedName: u.encodedName, properties: m}
}

func (u UnknownBlock) Properties() map[string]string {
	return u.properties
}
