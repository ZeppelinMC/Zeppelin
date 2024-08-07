package item

type ContainerLoot struct {
	LootTable string `nbt:"loot_table"`
	Seed      int32  `nbt:"seed"`
}
