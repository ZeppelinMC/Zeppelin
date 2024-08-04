package block

import (
	"strconv"
)

type Tripwire struct {
	East bool
	North bool
	Powered bool
	South bool
	West bool
	Attached bool
	Disarmed bool
}

func (b Tripwire) Encode() (string, BlockProperties) {
	return "minecraft:tripwire", BlockProperties{
		"east": strconv.FormatBool(b.East),
		"north": strconv.FormatBool(b.North),
		"powered": strconv.FormatBool(b.Powered),
		"south": strconv.FormatBool(b.South),
		"west": strconv.FormatBool(b.West),
		"attached": strconv.FormatBool(b.Attached),
		"disarmed": strconv.FormatBool(b.Disarmed),
	}
}

func (b Tripwire) New(props BlockProperties) Block {
	return Tripwire{
		North: props["north"] != "false",
		Powered: props["powered"] != "false",
		South: props["south"] != "false",
		West: props["west"] != "false",
		Attached: props["attached"] != "false",
		Disarmed: props["disarmed"] != "false",
		East: props["east"] != "false",
	}
}