package server

type Plugin struct {
	Identifier string
	OnLoad     func(*Server)
}
