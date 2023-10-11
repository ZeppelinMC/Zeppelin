package item

import (
	_ "embed"
	"encoding/json"
)

//go:embed items.json
var it []byte

type Item struct {
	ProtocolID int32 `json:"protocol_id"`
}

var items map[string]Item

func loadItems() {
	var itemData struct {
		Item struct {
			Default    string          `json:"default"`
			Entries    map[string]Item `json:"entries"`
			ProtocolID int32           `json:"protocol_id"`
		} `json:"minecraft:item"`
	}
	json.Unmarshal(it, &itemData)
	items = itemData.Item.Entries
}

func GetItem(name string) (item Item, ok bool) {
	if items == nil {
		loadItems()
	}
	it, ok := items[name]
	return it, ok
}
