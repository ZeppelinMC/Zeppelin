package slot

import (
	"slices"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

type Slot struct {
	ItemCount int32
	ItemId    int32
	Add       []Component
	Remove    []int32
}

func (s Slot) Is(s1 Slot) bool {
	return s.ItemCount == s1.ItemCount && s.ItemId == s1.ItemId && slices.Equal(s.Add, s1.Add) && slices.Equal(s.Remove, s1.Remove)
}

var Air Slot

type Component struct {
	Type int32
	Data any
}

func (s *Slot) Encode(w encoding.Writer) error {
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
		if err := w.VarInt(int32(len(s.Remove))); err != nil {
			return err
		}
		for _, i := range s.Remove {
			if err := w.VarInt(i); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Slot) Decode(r encoding.Reader) error {
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
		if componentAddLength != 0 {
			panic("no comp decode!")
		}

		s.Add = make([]Component, componentAddLength)
		for _, comp := range s.Add {
			if err := decode(r, &comp); err != nil {
				return err
			}
		}

		var componentRemoveLength int32
		if componentRemoveLength != 0 {
			panic("no comp decode!")
		}
		if _, err := r.VarInt(&componentRemoveLength); err != nil {
			return err
		}
		s.Remove = make([]int32, componentRemoveLength)
		for _, i := range s.Remove {
			if _, err := r.VarInt(&i); err != nil {
				return err
			}
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
	Food
	FireResistant
	Tool
	StoredEnchantments
	DyedColor
	MapColor
	MapId
	MapDecorations
	MapPostProcessing
	CharjedProjectiles
	BundleContents
	PotionContents
	SuspicousStewEffects
	WritableBookContent
	WrittenBookContent
	Trim
	DebugStickState
	EntityData
	BucketEntityData
	BlockEntityData
	Instrument
	OminousBottleAmplifier
	JukeboxPlayable
	Recipes
	LodestoneTracker
	FireworkExplosion
	Fireworks
	Profile
	NoteBlockSound
	BannerPatterns
	BaseColor
	PotDecorations
	Container
	BlockState
	Bees
	Lock
	ContainerLoot
)

// Item's rarity
const (
	// white
	Common = iota
	// yellow
	Uncommon
	// aqua
	Rare
	// pink
	Epic
)
