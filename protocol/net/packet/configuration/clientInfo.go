package configuration

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

const (
	ChatModeEnabled = iota
	ChatModeCommandsOnly
	ChatModeHidden
)

const (
	CapeEnabled = 1 << iota
	JacketEnabled
	LeftSleeveEnabled
	RightSleeveEnabled
	LeftPantsLegEnabled
	RightPantsLegEnabled
	HatEnabled
)

const (
	MainHandLeft = iota
	MainHandRight
)

// serverbound
const PacketIdClientInformation = 0x00

type ClientInformation struct {
	Locale              string
	ViewDistance        int8
	ChatMode            int32
	ChatColors          bool
	DisplayedSkinParts  byte
	MainHand            int32
	EnableTextFiltering bool
	AllowServerListing  bool
}

func (ClientInformation) ID() int32 {
	return 0x00
}

func (c *ClientInformation) Encode(w encoding.Writer) error {
	if err := w.String(c.Locale); err != nil {
		return err
	}
	if err := w.Byte(c.ViewDistance); err != nil {
		return err
	}
	if err := w.VarInt(c.ChatMode); err != nil {
		return err
	}
	if err := w.Bool(c.ChatColors); err != nil {
		return err
	}
	if err := w.Ubyte(c.DisplayedSkinParts); err != nil {
		return err
	}
	if err := w.VarInt(c.MainHand); err != nil {
		return err
	}
	if err := w.Bool(c.EnableTextFiltering); err != nil {
		return err
	}
	return w.Bool(c.AllowServerListing)
}

func (c *ClientInformation) Decode(r encoding.Reader) error {
	if err := r.String(&c.Locale); err != nil {
		return err
	}
	if err := r.Byte(&c.ViewDistance); err != nil {
		return err
	}
	if _, err := r.VarInt(&c.ChatMode); err != nil {
		return err
	}
	if err := r.Bool(&c.ChatColors); err != nil {
		return err
	}
	if err := r.Ubyte(&c.DisplayedSkinParts); err != nil {
		return err
	}
	if _, err := r.VarInt(&c.MainHand); err != nil {
		return err
	}
	if err := r.Bool(&c.EnableTextFiltering); err != nil {
		return err
	}
	return r.Bool(&c.AllowServerListing)
}
