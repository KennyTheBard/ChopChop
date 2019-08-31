package handle

import (
	"bufio"
	"fmt"
	"strings"

	cli "../cli"
	context "../context"
	component "../context/component"
)

func TaskFactory(task component.Action) func(*bufio.Scanner) {
	return func(scanner *bufio.Scanner) {
		target := task.NextWord()

		cli.Prompt()
		for scanner.Scan() {
			if strings.EqualFold(target, cli.GetInput(scanner)) {
				reward := task.GetReward()
				context.GlobalContext.GetInventory().AddItem(reward)
				fmt.Println("+1 " + reward)
			} else if strings.EqualFold("stop", cli.GetInput(scanner)) {
				break
			}

			target = task.NextWord()
			cli.Prompt()
		}
	}
}
