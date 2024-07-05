package configuration

import "aether/net/packet/login"

type CookieRequest struct{ login.CookieRequest }

func (CookieRequest) ID() int32 {
	return 0x00
}

type CookieResponse struct{ login.CookieResponse }

func (CookieResponse) ID() int32 {
	return 0x01
}
