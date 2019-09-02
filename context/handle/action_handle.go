package handle

import (
	"fmt"

	context ".."
	cli "../../cli"
	component "../component"
)

func ActionHandle() {
	curr := context.GlobalContext.GetCurrentLocation()
	wMap := context.GlobalContext.GetWorldMap()
	reg := context.GlobalContext.GetActionRegister()
	reader := context.GlobalContext.GetReader()
	reader.SetPrompt("action")

	for {
		if reader.IsInputEqual("what") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAvailableActions(),
				" * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			reader.Next()
			break

		} else if arg := reader.GetInput(); wMap.GetLocation(curr).HasAvailableAction(arg) {
			fmt.Println("You start to " + arg)
			actionFactory(reg.GetAction(arg))

		} else {
			fmt.Println("ERROR [" + arg + "]: This action cannot be done here!")
			break
		}
	}
}

func actionFactory(action component.Action) {
	reader := context.GlobalContext.GetReader()
	inventory := context.GlobalContext.GetInventory()
	reward := action.GetReward()

	for {
		target := action.NextWord()
		fmt.Println("Type '" + target + "'")

		if reader.IsInputEqual("stop") {
			reader.Next()
			break

		} else if reader.IsInputEqual(target) {
			reader.Next()
			inventory.AddItem(reward)
			fmt.Println(" +1 " + reward)
		}
	}
}
