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

		} else if reader.IsInputEqual("all") {
			items := inv.GetAllItems()

			if len(items) > 0 {
				fmt.Println(cli.BuildResponse(
					items,
					"In the inventory are:\n * ",
					"\n * ",
					""))
			} else {
				fmt.Println("The inventory is empty")
			}

		} else {
			item := reader.GetInput()
			fmt.Println(" " + item + ": " + strconv.Itoa(inv.GetItem(item)))
		}
	}
}
