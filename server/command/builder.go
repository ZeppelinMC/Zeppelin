package command

import "github.com/zeppelinmc/zeppelin/protocol/net/packet/play"

const (
	Bool = iota
	Float
	Double
	Integer
	Long
	String

	Entity
	GameProfile
	BlockPos
	ColumnPos
	Vec3
	Vec2
	BlockState
	BlockPredicate
	ItemStack
	ItemPredicate
	Color
	Component
	Style
	Message
	NBT
	NBTTag
	NBTPath
	Objective
	ObjectiveCriteria
	Operation
	Particle
	Angle
	Rotation
	ScoreboardSlot
	ScoreHolder
	Swizzle
	Team
	ItemSlot
	ResourceLocation
	Function
	EntityAnchor
	IntRange
	FloatRange
	Dimension
	Gamemode
	Time
	ResourceOrTag
	ResourceOrTagKey
	Resource
	ResourceKey
	TemplateMirror
	TemplateRotation
	Heightmap
	UUID
)

const (
	StringSingleWord = iota
	StringQuotablePhrase
	StringGreedyPhrase
)

type Node struct {
	play.Node
	children []Node
}

func (n *Node) Add(nodes ...Node) {
	n.children = append(n.children, nodes...)
}

func NewNode(n play.Node, children ...Node) Node {
	return Node{n, children}
}

func NewLiteral(name string, nodes ...Node) Node {
	return Node{
		Node: play.Node{
			Flags: play.NodeLiteral,
			Name:  name,
		},
		children: nodes,
	}
}

func NewBoolArgument(name string, nodes ...Node) Node {
	return Node{
		Node: play.Node{
			Flags:    play.NodeArgument,
			Name:     name,
			ParserId: Bool,
		},
		children: nodes,
	}
}

func NewIntegerArgument(name string, min, max *int32, nodes ...Node) Node {
	flags := byte(0)

	var m, x int32

	if min != nil {
		flags &= 0x01
		m = *min
	}
	if max != nil {
		flags &= 0x02
		x = *max
	}

	return Node{
		Node: play.Node{
			Flags:      play.NodeArgument,
			Name:       name,
			ParserId:   Integer,
			Properties: []any{flags, m, x},
		},
		children: nodes,
	}
}

func NewStringArgument(name string, typ int, nodes ...Node) Node {
	return Node{
		Node: play.Node{
			Flags:      play.NodeArgument,
			Name:       name,
			ParserId:   Integer,
			Properties: []any{typ},
		},
		children: nodes,
	}
}

func NewTimeArgument(name string, min int32, nodes ...Node) Node {
	return Node{
		Node: play.Node{
			Flags:      play.NodeArgument,
			Name:       name,
			ParserId:   Time,
			Properties: []any{min},
		},
		children: nodes,
	}
}
