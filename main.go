package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/nbt"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/network/registry"
)

func main() {
	var reg registry.DefaultRegistry
	file, _ := os.Open("server/network/registry.nbt")
	err := nbt.NewDecoder(file).Decode(&reg)
	fmt.Printf("%v %#v\n", err, reg)
	var logger = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)

	cfg, err := server.LoadConfig("config.toml")
	if err != nil {
		logger.Printf("%v. Using default config %#v", err, server.DefaultConfig)
		cfg = &server.DefaultConfig
	}

	st := minecraft.NewStatus(minecraft.Version{
		Protocol: 763,
		Text:     "1.20.1",
	}, cfg.MaxPlayers, cfg.Description, cfg.EnforceSecureChat, cfg.PreviewsChat)

	srv, err := cfg.NewServer(st)
	if err != nil {
		logger.Fatal(err)
	}
	srv.Logger = logger

	for {
		_, err := srv.Accept()
		if err != nil {
			logger.Fatal(err)
		}
	}
}
