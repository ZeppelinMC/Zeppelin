package block

import (
	"strconv"
)

type NoteBlock struct {
	Instrument string
	Note int
	Powered bool
}

func (b NoteBlock) Encode() (string, BlockProperties) {
	return "minecraft:note_block", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"instrument": b.Instrument,
		"note": strconv.Itoa(b.Note),
	}
}

func (b NoteBlock) New(props BlockProperties) Block {
	return NoteBlock{
		Note: atoi(props["note"]),
		Powered: props["powered"] != "false",
		Instrument: props["instrument"],
	}
}