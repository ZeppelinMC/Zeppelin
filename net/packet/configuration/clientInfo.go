package configuration

import "aether/net/io"

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

type ClientInformation struct {
	Locale              string
	ViewDistance        int8
	ChatMode            int32
	ChatColors          bool
	DisplayedSkinParts  byte
	EnableTextFiltering bool
	AllowServerListing  bool
}

func (ClientInformation) ID() int32 {
	return 0x00
}

func (c *ClientInformation) Encode(w io.Writer) error {
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
	if err := w.Bool(c.EnableTextFiltering); err != nil {
		return err
	}
	return w.Bool(c.AllowServerListing)
}

func (c *ClientInformation) Decode(r io.Reader) error {
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
	if err := r.Bool(&c.EnableTextFiltering); err != nil {
		return err
	}
	return r.Bool(&c.AllowServerListing)
}
