package handle

import (
	"fmt"
	"strings"

	cli "../cli"
	context "../context"
	component "../context/component"
)

func ActionHandle(args []string) {
	for _, arg := range args {
		curr := context.GlobalContext.GetCurrentLocation()
		wMap := context.GlobalContext.GetWorldMap()
		reg := context.GlobalContext.GetActionRegister()

		if strings.EqualFold(arg, "what") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAvailableActions(),
				" * ",
				"\n * ",
				""))

		} else {
			if wMap.GetLocation(curr).HasAvailableAction(arg) {
				fmt.Println("You start to " + arg)
				actionFactory(reg.GetAction(arg))

			} else {
				fmt.Println("ERROR [" + arg + "]: This action cannot be done here!")
				break
			}
		}
	}
}

func actionFactory(action component.Action) {
	scanner := context.GlobalContext.GetScanner()
	inventory := context.GlobalContext.GetInventory()
	reward := action.GetReward()

	for {
		target := action.NextWord()
		fmt.Println("Type '" + target + "'")

		cli.Prompt()
		scanner.Scan()

		if cli.GetInput(scanner) == "stop" {
			break

		} else if cli.GetInput(scanner) == target {
			inventory.AddItem(reward)
			fmt.Println(" +1 " + reward)
		}
	}
}
