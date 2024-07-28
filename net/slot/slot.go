package slot

import "github.com/zeppelinmc/zeppelin/net/io"

type Slot struct {
	ItemCount  int32
	ItemId     int32
	Components []Component
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
		if err := w.VarInt(int32(len(s.Components))); err != nil {
			return err
		}
		//TODO ENCODE COMPONENTS
		if err := w.VarInt(0); err != nil {
			return err
		}
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
