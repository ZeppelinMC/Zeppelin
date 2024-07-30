package block

import "strconv"

type BlockProperties = map[string]string

type Axis = string

const (
	AxisX Axis = "x"
	AxisY Axis = "y"
	AxisZ Axis = "z"
)

func atoi(str string) int {
	v, _ := strconv.Atoi(str)

	return v
}

type Block interface {
	// returns the name and properties of the block
	Encode() (name string, properties BlockProperties)
	// creates a new instance of the block with the specifiedproperties
	New(properties BlockProperties) Block
}

type UnknownBlock struct {
	name  string
	props BlockProperties
}

func (block UnknownBlock) Encode() (string, BlockProperties) {
	return block.name, block.props
}

func (block UnknownBlock) New(props BlockProperties) Block {
	return UnknownBlock{name: block.name, props: props}
}

// make sure unknown block implements block
var _ Block = (*UnknownBlock)(nil)
