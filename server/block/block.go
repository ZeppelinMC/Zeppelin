package block

type Block interface {

	// EncodedName is the string used to encode the block into disk
	EncodedName() string

	// New returns an instance of the block.
	// Passing nil maps are fine, since not all blocks have multiple states.
	New(map[string]string) Block

	Properties() map[string]string
}

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
