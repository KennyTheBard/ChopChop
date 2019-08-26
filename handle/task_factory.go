package handle

import (
	"bufio"
	"fmt"
	"strings"

	cli "../cli"
	context "../context"
)

func TaskFactory(task string, gContext context.GameContext) func(*bufio.Scanner) {
	return func(scanner *bufio.Scanner) {
		target := gContext.GetDictionary().NextWord(task)

		cli.Prompt()
		for scanner.Scan() {
			if strings.EqualFold(target, cli.GetInput(scanner)) {
				gContext.GetInventory().AddItem(task)
				fmt.Println("+1 " + task)
			} else if strings.EqualFold("stop", cli.GetInput(scanner)) {
				break
			}

			target = gContext.GetDictionary().NextWord(task)
			cli.Prompt()
		}
	}
}
