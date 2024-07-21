package play

import "github.com/zeppelinmc/zeppelin/net/packet/configuration"

//serverbound
const PacketIdClientInformation = 0x0A

type ClientInformation struct {
	configuration.ClientInformation
}

func (ClientInformation) ID() int32 {
	return 0x0A
}
