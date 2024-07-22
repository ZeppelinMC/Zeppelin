package std

import (
	_ "embed"

	"github.com/zeppelinmc/zeppelin/net/packet"
)

// Cached update tags packet

//go:embed tags.bin
var tagsBin []byte

var updateTags = packet.Raw(0x0D, tagsBin)
