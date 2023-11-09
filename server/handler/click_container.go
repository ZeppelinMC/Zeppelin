package handler

import (
	"fmt"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/inventory"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClickContainer(state *player.Player, pk *packet.ClickContainer) {
	if pk.WindowID == 0 {
		//fmt.Println(pk.Mode, pk.Button)
		switch pk.Mode {
		case 0:
			switch pk.Button {
			case 0:
				{
					if pk.Slot == -999 {
						// left click outside container
					} else {
						var slot *int16
						for s := range pk.Slots {
							slot = &s
							break
						}
						if slot == nil {
							fmt.Println("setting carried item")
							state.Inventory.SetCarriedItem(inventory.NetworkSlotToDataSlot(pk.Slot))
						} else {
							state.Inventory.Merge(inventory.NetworkSlotToDataSlot(pk.Slot))
						}
					}
				}
			case 1:
				{
					if pk.Slot == -999 {
						// right click outside container
					} else {
						fmt.Println("splitting")
						state.Inventory.Split(inventory.NetworkSlotToDataSlot(pk.Slot))
					}
				}
			}
		case 1:
			switch pk.Button {
			case 0, 1: // shift + left/right mouse click
			}
		case 2:
			switch pk.Button {
			case 40:
				{
					for s := range pk.Slots {
						state.Inventory.Swap(inventory.NetworkSlotToDataSlot(s), -106)
						break
					}
				}
			default:
				{
					if pk.Button < 9 {
						slot := inventory.NetworkSlotToDataSlot(pk.Slot)
						var new int8
						for s := range pk.Slots {
							new = inventory.NetworkSlotToDataSlot(s)
							break
						}
						state.Inventory.Swap(slot, new)
					}
				}
			}
		case 3:
			switch pk.Button {
			case 2: // middle click
			}
		case 4:
			switch pk.Button {
			case 0: // drop key (Q)
			case 1: // CTRL + drop key (Q)
			}
		case 5:
			switch pk.Button {
			case 0: // starting left mouse drag
			case 4: // starting right mouse drag
			case 8: // starting middle mouse drag

			case 1: // add slot for left mouse drag
			case 5: // add slot for right mouse drag
			case 9: // add slot for middle mouse drag

			case 2: // ending left mouse drag
			case 6: // ending right mouse drag
			case 10: // ending middle mouse drag
			}
		case 6:
			switch pk.Button {
			case 0:
				{
					state.Inventory.Collect(inventory.NetworkSlotToDataSlot(pk.Slot))
				}
			}
		}
	}
}
