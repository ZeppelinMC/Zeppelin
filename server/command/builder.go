package command

import "github.com/zeppelinmc/zeppelin/net/packet/play"

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

func NewCommand(name string, nodes ...Node) Node {
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
