package login

import (
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

type Textures struct {
	Timestamp         int64  `json:"timestamp"`
	ProfileId         string `json:"profileId"`
	ProfileName       string `json:"profileName"`
	SignatureRequired bool   `json:"signatureRequired"`
	Textures          struct {
		Skin struct {
			URL      string `json:"url"`
			Metadata struct {
				Model string `json:"model"`
			} `json:"metadata"`
		} `json:"SKIN"`
		Cape struct {
			URL string `json:"url"`
		} `json:"CAPE"`
	} `json:"textures"`
}

type Property struct {
	Name      string
	Value     string
	Signature string
}

// clientbound
const PacketIdLoginSuccess = 0x02

type LoginSuccess struct {
	UUID       uuid.UUID
	Username   string
	Properties []Property
}

func (LoginSuccess) ID() int32 {
	return PacketIdLoginSuccess
}

func (l *LoginSuccess) Encode(w encoding.Writer) error {
	if err := w.UUID(l.UUID); err != nil {
		return err
	}
	if err := w.String(l.Username); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(l.Properties))); err != nil {
		return err
	}
	for _, property := range l.Properties {
		if err := w.String(property.Name); err != nil {
			return err
		}
		if err := w.String(property.Value); err != nil {
			return err
		}
		if err := w.Bool(property.Signature != ""); err != nil {
			return err
		}
		if property.Signature != "" {
			if err := w.String(property.Signature); err != nil {
				return err
			}
		}
	}

	w.Bool(true) //TODO:remove for 1.21.3

	return nil
}

const PacketIdLoginAcknowledged = 0x03

type LoginAcknowledged struct{}

func (LoginAcknowledged) ID() int32 {
	return PacketIdLoginAcknowledged
}

func (*LoginAcknowledged) Decode(encoding.Reader) error {
	return nil
}
