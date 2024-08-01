package item

import (
	"fmt"

	"github.com/zeppelinmc/zeppelin/net/slot"
	"github.com/zeppelinmc/zeppelin/server/registry"
)

type DataSlot int8

func (slot DataSlot) Network() int32 {
	switch {
	case slot == 100:
		return 8
	case slot == 101:
		return 7
	case slot == 102:
		return 6
	case slot == 103:
		return 5
	case slot == -106:
		return 45
	case slot <= 8:
		return int32(slot + 36)
	case slot >= 80 && slot <= 83:
		return int32(slot - 79)
	default:
		return int32(slot)
	}
}

func dataSlot(network int32) DataSlot {
	switch {
	case network == 8:
		return 100
	case network == 7:
		return 101
	case network == 6:
		return 102
	case network == 5:
		return 103
	case network == 45:
		return -106
	case network >= 36 && network <= 44:
		return DataSlot(network - 36)
	case network >= 1 && network <= 4:
		return DataSlot(network + 79)
	default:
		return DataSlot(network)
	}
}

type Item struct {
	// The slot (as stored in the player data)
	Slot DataSlot `nbt:"Slot"`
	// the amount of items in the slot
	Count int32 `nbt:"count"`
	// The string id of this item
	Id string `nbt:"id"`
	// Components of this item (https://minecraft.wiki/w/Data_component_format#List_of_components)
	Components struct {
		AttributeModifiers       AttributeModifiers       `nbt:"minecraft:attribute_modifiers"`
		BannerPatterns           []BannerPattern          `nbt:"minecraft:banner_patterns"`
		BaseColor                BaseColor                `nbt:"minecraft:base_color"`
		Bees                     []Bee                    `nbt:"minecraft:bees"`
		BlockEntityData          BlockEntityData          `nbt:"minecraft:block_entity_data"`
		BlockState               BlockState               `nbt:"minecraft:block_state"`
		BucketEntityData         BucketEntityData         `nbt:"minecraft:bucket_entity_data"`
		BundleContents           []BundleContent          `nbt:"minecraft_bundle_contents"`
		CanBreak                 CanBreak                 `nbt:"minecraft:can_break"`
		ChargedProjectiles       []ChargedProjectile      `nbt:"minecraft:charged_projectiles"`
		Container                []Container              `nbt:"minecraft:container"`
		ContainerLoot            ContainerLoot            `nbt:"minecraft_container_loot"`
		CustomData               CustomData               `nbt:"minecraft:custom_data"`
		CustomModelData          CustomModelData          `nbt:"minecraft:custom_model_data"`
		CustomName               CustomName               `nbt:"minecraft:custom_name"`
		Damage                   Damage                   `nbt:"minecraft:damage"`
		DebugStickState          DebugStickState          `nbt:"minecraft:debug_stick_state"`
		DyedColor                DyedColor                `nbt:"minecraft:dyed_color"`
		EnchantmentGlintOverride EnchantmentGlintOverride `nbt:"minecraft:enchantment_glint_override"`
		Enchantments             Enchantments             `nbt:"minecraft:enchantments"`
		EntityData               EntityData               `nbt:"minecraft:entity_data"`
		FireResistant            FireResistant            `nbt:"minecraft:fire_resistant"`
		FireworkExplosion        FireworkExplosion        `nbt:"minecraft:firework_explosion"`
		Fireworks                Fireworks                `nbt:"minecraft:fireworks"`
		Food                     Food                     `nbt:"minecraft:food"`
		HideAdditionalTooltip    HideAdditionalTooltip    `nbt:"minecraft:hide_additional_tooltip"`
		HideTooltip              HideTooltip              `nbt:"minecraft:hide_tooltip"`
		Instrument               Instrument               `nbt:"minecraft:instrument"`
		IntangibleProjectile     IntangibleProjectile     `nbt:"minecraft:intangible_projectile`
		ItemName                 ItemName                 `nbt:"minecraft:item_name"`
		JukeboxPlayable          JukeboxPlayable          `nbt:"minecraft:jukebox_playable"`
		Lock                     Lock                     `nbt:"minecraft:lock"`
		LodestoneTracker         LodestoneTracker         `nbt:"minecraft:lodestone_tracker"`
		Lore                     []Lore                   `nbt:"minecraft:lore"`
		MapColor                 MapColor                 `nbt:"minecraft:map_color"`
		MapDecorations           MapDecorations           `nbt:"minecraft:map_decorations"`
		MapID                    MapID                    `nbt:"minecraft:map_id"`
		MaxDamage                MaxDamage                `nbt:"minecraft:max_damage"`
		MaxStackSize             MaxStackSize             `nbt:"minecraft:max_stack_size"`
		NoteBlockSound           NoteBlockSound           `nbt:"minecraft:note_block_sound"`
		OminousBottleAmplifier   OminousBottleAmplifier   `nbt:"minecraft:ominous_bottle_amplifier"`
		PotDecorations           []PotDecoration          `nbt:"minecraft:pot_decorations"`
		PotionContents           PotionContents           `nbt:"minecraft:potion_contents"`
		Profile                  Profile                  `nbt:"minecraft:profile"`
		Rarity                   Rarity                   `nbt:"minecraft:rarity"`
		Recipes                  []Recipe                 `nbt:"minecraft:recipes`
		RepairCost               RepairCost               `nbt:"minecraft:repair_cost"`
		StoredEnchantments       StoredEnchantments       `nbt:"minecraft:stored_enchantments"`
		SuspiciousStewEffects    []SuspiciousStewEffect   `nbt:"minecraft:suspicious_stew_effects"`
		Tool                     Tool                     `nbt:"minecraft:tool"`
		Trim                     Trim                     `nbt:"minecraft:trim"`
		Unbreakable              Unbreakable              `nbt:"minecraft:unbreakable"`
		WritableBookContent      WritableBookContent      `nbt:"minecraft:writable_book_content"`
		WrittenBookContent       WrittenBookContent       `nbt:"minecraft:written_book_content"`
		CreativeSlotLock         CreativeSlotLock         `nbt:"minecraft:creative_slot_lock"`
		MapPostProcessing        MapPostProcessing        `nbt:"minecraft:map_post_processing"`
	} `nbt:"components"`
}

// New creates an item from the slot provided
func New(slot int32, item slot.Slot) (Item, error) {
	i := Item{
		Slot:  dataSlot(slot),
		Count: item.ItemCount,
	}
	id, ok := registry.Item.NameOf(item.ItemId)
	if !ok {
		return i, fmt.Errorf("invalid item id")
	}

	i.Id = id

	return i, nil
}
