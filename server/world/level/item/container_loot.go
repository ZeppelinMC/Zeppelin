package item

type ContainerLoot struct {
	LootTable string `nbt:"loot_table"`
	Seed      int64  `nbt:"seed"`
}
