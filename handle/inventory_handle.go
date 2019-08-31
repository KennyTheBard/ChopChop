package handle

import (
	"fmt"
	"strconv"

	context "../context"
)

func InventoryHandle(args []string) {
	for _, arg := range args {
		inv := context.GlobalContext.GetInventory()

		fmt.Println(" " + arg + ": " + strconv.Itoa(inv.GetItem(arg)))
	}
}
