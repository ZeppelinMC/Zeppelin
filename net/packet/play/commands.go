package play

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/net/io"
)

type Node struct {
	Flags        int8
	Children     []int32
	RedirectNode int32
	Name         string
	ParserId     int32

	/*
		float64 = Double
		float32 = Float
		int8 = Byte
		int32 = Int
		int64 = Long
		string = Identifier
		int = VarInt
	*/
	Properties      []any
	SuggestionsType string
}

const NodeRoot = iota

const (
	NodeLiteral = 1 << iota
	NodeArgument

	NodeExecutable
	NodeRedirect
	NodeHasSuggestionsType
)

const NodeType = 0x03

// clientbound
const PacketIdCommands = 0x11

type Commands struct {
	Nodes     []Node
	RootIndex int32
}

func (Commands) ID() int32 {
	return PacketIdCommands
}

func (c *Commands) Encode(w io.Writer) error {
	if err := w.VarInt(int32(len(c.Nodes))); err != nil {
		return err
	}
	for _, node := range c.Nodes {
		if err := c.encodeNode(w, node); err != nil {
			return err
		}
	}
	return w.VarInt(c.RootIndex)
}

func (c *Commands) encodeNode(w io.Writer, node Node) error {
	if err := w.Byte(node.Flags); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(node.Children))); err != nil {
		return err
	}
	for _, child := range node.Children {
		if err := w.VarInt(child); err != nil {
			return err
		}
	}
	if node.Flags&NodeRedirect != 0 {
		if err := w.VarInt(node.RedirectNode); err != nil {
			return err
		}
	}
	if node.Flags&NodeType > 0 {
		if err := w.String(node.Name); err != nil {
			return err
		}
	}
	if node.Flags&NodeArgument != 0 {
		if err := w.VarInt(node.ParserId); err != nil {
			return err
		}
	}
	if node.Flags&NodeArgument != 0 {
		for _, p := range node.Properties {
			switch prop := p.(type) {
			case float64:
				if err := w.Double(prop); err != nil {
					return err
				}
			case float32:
				if err := w.Float(prop); err != nil {
					return err
				}
			case int8:
				if err := w.Byte(prop); err != nil {
					return err
				}
			case int32:
				if err := w.Int(prop); err != nil {
					return err
				}
			case int64:
				if err := w.Long(prop); err != nil {
					return err
				}
			case string:
				if err := w.Identifier(prop); err != nil {
					return err
				}
			case int:
				if err := w.VarInt(int32(prop)); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unrecognized node property type %T", p)
			}
		}
	}
	if node.Flags&NodeHasSuggestionsType != 0 {
		if err := w.Identifier(node.SuggestionsType); err != nil {
			return err
		}
	}
	return nil
}

func (*Commands) Decode(io.Reader) error {
	return nil //TODO
}
