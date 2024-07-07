package server

import (
	"aether/net"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"aether/net/registry"
	"fmt"
)

type Server struct {
	cfg      ServerConfig
	listener *net.Listener
	ticker   Ticker
}

func (srv Server) Start() {
	fmt.Println("started ticker!")
	srv.ticker.Start()
	fmt.Println("started server")
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			fmt.Println("server error", err)
			return
		}
		if conn == nil {
			continue
		}
		fmt.Println("new connection from player", conn.Username())
		for _, packet := range registry.RegistryMap.Packets() {
			fmt.Println(conn.WritePacket(packet))
		}
		conn.WritePacket(configuration.FinishConfiguration{})
		conn.SetState(net.PlayState)
		conn.WritePacket(&play.Login{
			EntityID: 1,

			ViewDistance:        12,
			SimulationDistance:  12,
			EnableRespawnScreen: true,
			DimensionType:       0,
			DimensionName:       "minecraft:overworld",
			GameMode:            1,
		})
	}
}
