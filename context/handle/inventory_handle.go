package handle

import (
	"fmt"
	"strconv"

	context ".."
	cli "../../cli"
)

func InventoryHandle() {
	inv := context.GlobalContext.GetInventory()
	reader := context.GlobalContext.GetReader()
	cli.SetPrompt("inventory")

	for {
		if reader.IsInputEqual("back") {
			break

		} else {
			item := reader.GetInput()
			fmt.Println(" " + item + ": " + strconv.Itoa(inv.GetItem(item)))
		}
	}
}
