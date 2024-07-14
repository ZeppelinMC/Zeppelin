package play

import (
	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet/login"
	"github.com/google/uuid"
)

const (
	ActionAddPlayer = 1 << iota
	ActionInitializeChat
	ActionUpdateGameMode
	ActionUpdateListed
	ActionUpdateLatency
	ActionUpdateDisplayName
)

type PlayerAction struct {
	// Add Player
	Name       string
	Properties []login.Property
	// Initialize Chat
	HasSignatureData bool
	Session          PlayerSession
	// Update Game Mode
	GameMode int32
	// Update Listed
	Listed bool
	// Update Latency
	Ping int32
	// Update Display Name
	HasDisplayName bool
	DisplayName    chat.TextComponent
}

// clientbound
const PacketIdPlayerInfoUpdate = 0x3E

type PlayerInfoUpdate struct {
	Actions int8
	Players map[uuid.UUID]PlayerAction
}

func (PlayerInfoUpdate) ID() int32 {
	return 0x3E
}

func (p *PlayerInfoUpdate) Encode(w io.Writer) error {
	if err := w.Byte(p.Actions); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(p.Players))); err != nil {
		return err
	}
	for uuid, player := range p.Players {
		if err := w.UUID(uuid); err != nil {
			return err
		}
		if p.Actions&ActionAddPlayer != 0 {
			if err := w.String(player.Name); err != nil {
				return err
			}
			if err := w.VarInt(int32(len(player.Properties))); err != nil {
				return err
			}
			for _, property := range player.Properties {
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
		}
		if p.Actions&ActionInitializeChat != 0 {
			if err := w.Bool(player.HasSignatureData); err != nil {
				return err
			}
			if player.HasSignatureData {
				if err := w.UUID(player.Session.SessionID); err != nil {
					return err
				}
				if err := w.Long(player.Session.ExpiresAt); err != nil {
					return err
				}
				if err := w.ByteArray(player.Session.PublicKey); err != nil {
					return err
				}
				if err := w.ByteArray(player.Session.KeySignature); err != nil {
					return err
				}
			}
		}
		if p.Actions&ActionUpdateGameMode != 0 {
			if err := w.VarInt(player.GameMode); err != nil {
				return err
			}
		}
		if p.Actions&ActionUpdateListed != 0 {
			if err := w.Bool(player.Listed); err != nil {
				return err
			}
		}
		if p.Actions&ActionUpdateLatency != 0 {
			if err := w.VarInt(player.Ping); err != nil {
				return err
			}
		}
		if p.Actions&ActionUpdateDisplayName != 0 {
			if err := w.Bool(player.HasDisplayName); err != nil {
				return err
			}
			return w.TextComponent(player.DisplayName)
		}
	}
	return nil
}

func (p *PlayerInfoUpdate) Decode(r io.Reader) error {
	return nil //TODO
}
