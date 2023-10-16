package handlers

import (
	"github.com/aimjel/minecraft/packet"
)

func ClientSettings(controller Controller, pk *packet.ClientSettings) {
	controller.SetClientSettings(pk)
	controller.BroadcastSkinData()
}
