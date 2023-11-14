package pos

import (
	"fmt"

	"github.com/aimjel/minecraft/protocol/types"
)

type BlockPosition [3]int64

func (b BlockPosition) X() int64 {
	return b[0]
}

func (b BlockPosition) Y() int64 {
	return b[1]
}

func (b BlockPosition) Z() int64 {
	return b[2]
}

func (b BlockPosition) Add(b1 BlockPosition) BlockPosition {
	return [3]int64{
		b.X() + b1.X(),
		b.Y() + b1.Y(),
		b.Z() + b1.Z(),
	}
}

func (b BlockPosition) String() string {
	return fmt.Sprintf("(%d %d %d)", b.X(), b.Y(), b.Z())
}

func (b BlockPosition) Data() types.Position {
	return types.Position(((b.X() & 0x3FFFFFF) << 38) | ((b.Z() & 0x3FFFFFF) << 12) | (b.Y() & 0xFFF))
}
