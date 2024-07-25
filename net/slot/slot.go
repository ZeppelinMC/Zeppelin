package slot

type Slot struct {
	ItemCount  int32
	ItemId     int32
	Components []any
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
