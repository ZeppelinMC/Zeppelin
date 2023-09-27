package server

import "fmt"

type ConsoleExecutor struct {
	Server *Server
}

func (p *ConsoleExecutor) SystemChatMessage(s string) error {
	fmt.Println(s)
	return nil
}
