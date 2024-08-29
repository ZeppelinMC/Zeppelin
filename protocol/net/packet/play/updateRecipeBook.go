package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

const (
	UpdateRecipeBookActionInit = iota
	UpdateRecipeBookActionAdd
	UpdateRecipeBookActionRemove
)

// clientbound
const PacketIdUpdateRecipeBook = 0x41

type UpdateRecipeBook struct {
	Action int32

	CraftingRecipeBookOpen         bool
	CraftingRecipeBookFilterActive bool

	SmeltingRecipeBookOpen         bool
	SmeltingRecipeBookFilterActive bool

	BlastFurnaceRecipeBookOpen         bool
	BlastFurnaceRecipeBookFilterActive bool

	SmokerRecipeBookOpen         bool
	SmokerRecipeBookFilterActive bool

	Array1 []string //init: to be displayed, add/rem: recipes
	Array2 []string //init: recipes, add/rem: unused
}

func (UpdateRecipeBook) ID() int32 {
	return PacketIdUpdateRecipeBook
}

func (u *UpdateRecipeBook) Encode(w io.Writer) error {
	if err := w.VarInt(u.Action); err != nil {
		return err
	}
	if err := w.Bool(u.CraftingRecipeBookOpen); err != nil {
		return err
	}
	if err := w.Bool(u.CraftingRecipeBookFilterActive); err != nil {
		return err
	}

	if err := w.Bool(u.SmeltingRecipeBookOpen); err != nil {
		return err
	}
	if err := w.Bool(u.SmeltingRecipeBookFilterActive); err != nil {
		return err
	}

	if err := w.Bool(u.BlastFurnaceRecipeBookOpen); err != nil {
		return err
	}
	if err := w.Bool(u.BlastFurnaceRecipeBookFilterActive); err != nil {
		return err
	}

	if err := w.Bool(u.SmokerRecipeBookOpen); err != nil {
		return err
	}
	if err := w.Bool(u.SmokerRecipeBookFilterActive); err != nil {
		return err
	}

	if err := w.VarInt(int32(len(u.Array1))); err != nil {
		return err
	}
	for _, str := range u.Array1 {
		if err := w.Identifier(str); err != nil {
			return err
		}
	}
	if u.Action == UpdateRecipeBookActionInit {
		if err := w.VarInt(int32(len(u.Array2))); err != nil {
			return err
		}
		for _, str := range u.Array2 {
			if err := w.Identifier(str); err != nil {
				return err
			}
		}
	}

	return nil
}

func (*UpdateRecipeBook) Decode(io.Reader) error {
	return nil //TODO
}
