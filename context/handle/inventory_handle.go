package handle

import (
	"fmt"
	"strconv"

	context ".."
)

func InventoryHandle() {
	inv := context.GlobalContext.GetInventory()
	reader := context.GlobalContext.GetReader()

	for {
		if reader.IsInputEqual("stop") {
			reader.Next()
			break

		} else {
			item := reader.GetInput()
			fmt.Println(" " + item + ": " + strconv.Itoa(inv.GetItem(item)))
		}
	}
}
