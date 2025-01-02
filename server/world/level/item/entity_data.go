package item

import "github.com/zeppelinmc/zeppelin/server/entity"

type EntityData struct {
	EntityData entity.LevelEntity `nbt:"minecraft:entity_data"` // NBT
}
