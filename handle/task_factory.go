package handle

import (
	"bufio"
	"fmt"
	"strings"

	cli "../cli"
	context "../context"
)

var GlobalContext context.GameContext

func TaskFactory(task string) func(*bufio.Scanner) {
	return func(scanner *bufio.Scanner) {
		target := GlobalContext.GetDictionary().NextWord(task)

		cli.Prompt()
		for scanner.Scan() {
			if strings.EqualFold(target, cli.GetInput(scanner)) {
				GlobalContext.GetInventory().AddItem(task)
				fmt.Println("+1 " + task)
			} else if strings.EqualFold("stop", cli.GetInput(scanner)) {
				break
			}

			target = GlobalContext.GetDictionary().NextWord(task)
			cli.Prompt()
		}
	}
}
