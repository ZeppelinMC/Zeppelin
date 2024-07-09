package play

import "aether/net/packet/configuration"

type ClientInformation struct {
	configuration.ClientInformation
}

func (ClientInformation) ID() int32 {
	return 0x0A
}
