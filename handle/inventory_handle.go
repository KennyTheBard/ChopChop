package handle

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	cli "../cli"
	context "../context"
)

func InventoryHandler(gContext context.GameContext) func(*bufio.Scanner) {
	return func(scanner *bufio.Scanner) {
		fmt.Print("What are you looking for in your inventory?\n>> ")

		for scanner.Scan() {
			if strings.EqualFold("stop", cli.GetInput(scanner)) {
				break
			} else {
				if item, ok := gContext.GetInventory()[cli.GetInput(scanner)]; ok {
					fmt.Println(cli.GetInput(scanner) + ": " + strconv.Itoa(item))
				} else {
					fmt.Println("None")
				}
			}

			cli.Prompt()
		}
	}

}
