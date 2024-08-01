package item

type Bee *struct {
	EntityData         int32 `nbt:"entity_data"`
	MinimumTicksInHive int32 `nbt:"min_ticks_in_hive"`
	TicksInHive        int32 `nbt:"ticks_in_hive"`
}
