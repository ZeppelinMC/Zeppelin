package slot

import "github.com/zeppelinmc/zeppelin/net/io"

type Slot struct {
	ItemCount   int32
	ItemId      int32
	Add, Remove []Component
}

type Component struct {
	Type int32
	Data any
}

func (s *Slot) Encode(w io.Writer) error {
	if err := w.VarInt(s.ItemCount); err != nil {
		return err
	}
	if s.ItemCount > 0 {
		if err := w.VarInt(s.ItemId); err != nil {
			return err
		}
		if err := w.VarInt(int32(len(s.Add))); err != nil {
			return err
		}
		//TODO ENCODE COMPONENTS
		if err := w.VarInt(0); err != nil {
			return err
		}
	}
	return nil
}

func (s *Slot) Decode(r io.Reader) error {
	if _, err := r.VarInt(&s.ItemCount); err != nil {
		return err
	}
	if s.ItemCount > 0 {
		if _, err := r.VarInt(&s.ItemId); err != nil {
			return err
		}
		var componentAddLength int32
		if _, err := r.VarInt(&componentAddLength); err != nil {
			return err
		}
		s.Add = make([]Component, componentAddLength)
		//TODO DECODE COMPONENTS

		var componentRemoveLength int32
		if _, err := r.VarInt(&componentRemoveLength); err != nil {
			return err
		}
		s.Remove = make([]Component, componentRemoveLength)
	}
	return nil
}

const (
	CustomData = iota
	MaxStackSize
	MaxDamage
	Damage
	Unbreakable
	CustomName
	ItemName
	Lore
	Rarity
	Enchantments
	CanPlaceOn
	CanBreak
	AttributeModifiers
	CustomModelData
	HideAdditionalTooltip
	HideTooltip
	RepairCost
	CreativeSlotLock
	EnchantmentGlintOverride
	IntangibleProjectile
)
