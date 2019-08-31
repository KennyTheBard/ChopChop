package handle

import (
	"fmt"
	"strconv"
	"strings"

	cli "../cli"
	context "../context"
)

func InventoryHandler(args []string) {
	fmt.Print("What are you looking for in your inventory?")

	cli.Prompt()
	for scanner.Scan() {
		if strings.EqualFold("stop", cli.GetInput(scanner)) {
			break
		} else {
			if item, ok := context.GlobalContext.GetInventory()[cli.GetInput(scanner)]; ok {
				fmt.Println(cli.GetInput(scanner) + ": " + strconv.Itoa(item))
			} else {
				fmt.Println("None")
			}
		}

		cli.Prompt()
	}
}
