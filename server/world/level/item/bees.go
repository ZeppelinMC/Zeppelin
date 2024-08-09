package item

import "github.com/zeppelinmc/zeppelin/server/entity"

type Bee struct {
	EntityData         entity.LevelEntity `nbt:"entity_data"` // import cycle :(
	MinimumTicksInHive int32              `nbt:"min_ticks_in_hive"`
	TicksInHive        int32              `nbt:"ticks_in_hive"`
}
