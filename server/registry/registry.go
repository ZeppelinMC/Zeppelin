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

var items map[string]Item
var entities map[string]Item

type registry struct {
	Default    string          `nbt:"default"`
	Entries    map[string]Item `nbt:"entries"`
	ProtocolID int32           `nbt:"protocol_id"`
}

func loadregistry() {
	var data struct {
		Item       registry `nbt:"minecraft:item"`
		EntityType registry `nbt:"minecraft:entity_type"`
	}
	nbt.Unmarshal(rg, &data)
	items = data.Item.Entries
	entities = data.EntityType.Entries
}

func GetItem(name string) (item Item, ok bool) {
	if items == nil {
		loadregistry()
	}
	it, ok := items[name]
	return it, ok
}

func GetEntity(name string) (item Item, ok bool) {
	if entities == nil {
		loadregistry()
	}
	it, ok := entities[name]
	return it, ok
}
