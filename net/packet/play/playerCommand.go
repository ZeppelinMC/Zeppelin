package play

import "github.com/dynamitemc/aether/net/io"

// serverbound
const PacketIdPlayerCommand = 0x25

const (
	ActionIdStartSneaking = iota
	ActionIdStopSneaking
	ActionIdLeaveBed
	ActionIdStartSprinting
	ActionIdStopSprinting
	ActionIdStartJumpWithHorse
	ActionIdStopJumpWithHorse
	ActionIdOpenVehicleInventory
	ActionIdStartFlyingWithElytra
)

type PlayerCommand struct {
	EntityId  int32
	ActionId  int32
	JumpBoost int32
}

func (PlayerCommand) ID() int32 {
	return PacketIdPlayerCommand
}

func (p *PlayerCommand) Encode(w io.Writer) error {
	if err := w.VarInt(p.EntityId); err != nil {
		return err
	}
	if err := w.VarInt(p.ActionId); err != nil {
		return err
	}
	return w.VarInt(p.JumpBoost)
}

func (p *PlayerCommand) Decode(r io.Reader) error {
	if _, err := r.VarInt(&p.EntityId); err != nil {
		return err
	}
	if _, err := r.VarInt(&p.ActionId); err != nil {
		return err
	}
	_, err := r.VarInt(&p.JumpBoost)
	return err
}
