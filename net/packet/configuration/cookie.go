package configuration

import "github.com/zeppelinmc/zeppelin/net/packet/login"

//clientbound
const PacketIdCookieRequest = 0x00

type CookieRequest struct{ login.CookieRequest }

func (CookieRequest) ID() int32 {
	return 0x00
}

//serverbound
const PacketIdCookieResponse = 0x01

type CookieResponse struct{ login.CookieResponse }

func (CookieResponse) ID() int32 {
	return 0x01
}
