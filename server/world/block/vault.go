package block

import (
	"strconv"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
)

type Vault struct {
	Facing string
	Ominous bool
	VaultState string
}

func (b Vault) Encode() (string, BlockProperties) {
	return "minecraft:vault", BlockProperties{
		"facing": b.Facing,
		"ominous": strconv.FormatBool(b.Ominous),
		"vault_state": b.VaultState,
	}
}

func (b Vault) New(props BlockProperties) Block {
	return Vault{
		Facing: props["facing"],
		Ominous: props["ominous"] != "false",
		VaultState: props["vault_state"],
	}
}

func (b Vault) BlockEntity(pos pos.BlockPosition) chunk.BlockEntity {
	return chunk.BlockEntity{
		Id:    "minecraft:vault",
		X:     pos.X(), Y: pos.Y(), Z: pos.Z(),
	}
}