package registry

type Element struct {
	Element map[string]any `nbt:"element"`
	ID      int            `nbt:"id"`
	Name    string         `nbt:"name"`
}

type Registry struct {
	Type  string    `nbt:"type"`
	Value []Element `nbt:"value"`
}

func (r Registry) Lookup(name string) (e Element, ok bool) {
	for _, el := range r.Value {
		if el.Name == name {
			return el, true
		}
	}
	return e, false
}

type DefaultRegistry struct {
	ChatType      Registry `nbt:"minecraft:chat_type"`
	DamageType    Registry `nbt:"minecraft:damage_type"`
	DimensionType Registry `nbt:"minecraft:dimension_type"`
	TrimMaterial  Registry `nbt:"minecraft:trim_material"`
	TrimPattern   Registry `nbt:"minecraft:trim_pattern"`
	Biome         Registry `nbt:"minecraft:worldgen/biome"`
}
