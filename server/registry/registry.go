package registry

import (
	_ "embed"
	"encoding/json"
)

//go:embed registries.json
var rg []byte

type Item struct {
	ProtocolID int32 `json:"protocol_id"`
}

var items map[string]Item
var entities map[string]Item

type registry struct {
	Default    string          `json:"default"`
	Entries    map[string]Item `json:"entries"`
	ProtocolID int32           `json:"protocol_id"`
}

func loadregistry() {
	var data struct {
		Item       registry `json:"minecraft:item"`
		EntityType registry `json:"minecraft:entity_type"`
	}
	json.Unmarshal(rg, &data)
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
