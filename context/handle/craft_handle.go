package handle

import (
	"fmt"

	context ".."
	cli "../../cli"
	component "../component"
)

func CraftHandle() {
	reg := context.GlobalContext.GetBlueprintRegister()
	reader := context.GlobalContext.GetReader()

	for {
		cli.SetPrompt("craft")

		if reader.IsInputEqual("help") {
			fmt.Println(cli.BuildResponse(
				reg.GetAllBlueprints(),
				"Available blueprints are:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else if arg := reader.GetInput(); reg.HasBlueprint(arg) {
			if craftItem(reg.GetBlueprint(arg)) {
				fmt.Println("You crafted 1 " + arg)
			} else {
				fmt.Println("Not enough resources to craft " + arg)
			}

		} else {
			fmt.Println("ERROR [" + arg + "]: This action cannot be done here!")
		}
	}
}

func craftItem(bp component.Blueprint) bool {
	inventory := context.GlobalContext.GetInventory()

	for item, cost := range bp.GetAllCosts() {
		if !inventory.HasItemAtLeast(item, cost) {
			return false
		}
	}

	for item, cost := range bp.GetAllCosts() {
		inventory.TakeItems(item, cost)
		inventory.AddItems(item, 1)
	}
	return true
}
