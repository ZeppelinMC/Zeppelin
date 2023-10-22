package registry

import (
	_ "embed"

	"github.com/aimjel/minecraft/nbt"
)

//go:embed registries.nbt
var rg []byte

type Item struct {
	ProtocolID int32 `nbt:"protocol_id"`
}

var data struct {
	SoundEvent registry `nbt:"minecraft:sound_event"`
	Item       registry `nbt:"minecraft:item"`
	EntityType registry `nbt:"minecraft:entity_type"`
}

type registry struct {
	Default    string          `nbt:"default"`
	Entries    map[string]Item `nbt:"entries"`
	ProtocolID int32           `nbt:"protocol_id"`
}

func loadregistry() {
	nbt.Unmarshal(rg, &data)
}

func GetItem(name string) (item Item, ok bool) {
	if data.Item.Entries == nil {
		loadregistry()
	}
	it, ok := data.Item.Entries[name]
	return it, ok
}

func GetEntity(name string) (item Item, ok bool) {
	if data.EntityType.Entries == nil {
		loadregistry()
	}
	it, ok := data.EntityType.Entries[name]
	return it, ok
}

func GetSound(name string) (item Item, ok bool) {
	if data.SoundEvent.Entries == nil {
		loadregistry()
	}
	it, ok := data.SoundEvent.Entries[name]
	return it, ok
}
