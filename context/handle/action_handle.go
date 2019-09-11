package handle

import (
	"fmt"
	"strconv"

	context ".."
	cli "../../cli"
	component "../component"
)

func ActionHandle() {
	curr := context.GlobalContext.GetCurrentLocation()
	wMap := context.GlobalContext.GetWorldMap()
	reg := context.GlobalContext.GetActionRegister()
	reader := context.GlobalContext.GetReader()

	for {
		cli.SetPrompt("action")

		if reader.IsInputEqual("help") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAvailableActions(),
				"Available actions are:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else if arg := reader.GetInput(); wMap.GetLocation(curr).HasAvailableAction(arg) {
			fmt.Println("You start to " + arg)
			actionFactory(reg.GetAction(arg))

		} else {
			fmt.Println("ERROR [" + arg + "]: This action cannot be done here!")
		}
	}
}

func actionFactory(action component.Action) {
	reader := context.GlobalContext.GetReader()
	inventory := context.GlobalContext.GetInventory()
	reward := action.GetReward()

	rewardQuantity := 1

	if inventory.HasItem(action.GetTool()) {
		rewardQuantity = 2
		inventory.TakeItems(action.GetTool(), 1)
	}

	for {
		cli.SetPrompt("action - " + action.GetName())

		target := action.NextWord()
		fmt.Println("Type '" + target + "'")

		if reader.IsInputEqual("stop") {
			break

		} else if reader.IsInputEqual(target) {
			inventory.AddItems(reward, rewardQuantity)
			fmt.Println(" +" + strconv.Itoa(rewardQuantity) + " " + reward)

		} else {
			fmt.Println(" " + reader.GetInput() + " different from " + target)
		}
	}
}
