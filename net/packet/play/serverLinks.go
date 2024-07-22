package play

import (
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/text"
)

// clientbound
const PacketIdServerLinks = 0x7B

const (
	LabelBugRepot = iota
	LabelCommunityGuidelines
	LabelSupport
	LabelStatus
	LabelFeedback
	LabelCommunity
	LabelWebsite
	LabelForums
	LabelNews
	LabelAnnouncements
)

type Link struct {
	BuiltIn bool

	BuiltInLabel int32              // used if built in is true
	Label        text.TextComponent // used if built in is false

	URL string
}

type ServerLinks struct {
	Links []Link
}

func (ServerLinks) ID() int32 {
	return PacketIdServerLinks
}

func (s *ServerLinks) Encode(w io.Writer) error {
	if err := w.VarInt(int32(len(s.Links))); err != nil {
		return err
	}
	for _, link := range s.Links {
		if err := w.Bool(link.BuiltIn); err != nil {
			return err
		}
		if link.BuiltIn {
			if err := w.VarInt(link.BuiltInLabel); err != nil {
				return err
			}
		} else {
			if err := w.TextComponent(link.Label); err != nil {
				return err
			}
		}
		if err := w.String(link.URL); err != nil {
			return err
		}
	}
	return nil
}

func (s *ServerLinks) Decode(r io.Reader) error {
	return nil //TODO
}
