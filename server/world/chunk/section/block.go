package section

type Block interface {
	// returns the name and properties of the block
	Encode() (name string, properties map[string]string)
	// creates a new instance of the block with the specified properties
	New(properties map[string]string) Block
}

type UnknownBlock struct {
	name  string
	props map[string]string
}

func (block UnknownBlock) Encode() (string, map[string]string) {
	return block.name, block.props
}

func (block UnknownBlock) New(props map[string]string) Block {
	return UnknownBlock{name: block.name, props: props}
}

// make sure unknown block implements block
var _ Block = (*UnknownBlock)(nil)
