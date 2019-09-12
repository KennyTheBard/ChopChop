package handle

import (
	"fmt"

	context ".."
	cli "../../cli"
)

func TravelHandle() {
	curr := context.GlobalContext.GetCurrentLocation()
	wMap := context.GlobalContext.GetWorldMap()
	reader := context.GlobalContext.GetReader()

	for {
		SetPrompt("travel")

		if reader.IsInputEqual("help") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAdjacentLocations(),
				"Close locations are:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else if arg := reader.GetInput(); wMap.GetLocation(curr).HasAdjacentLocation(arg) {
			context.GlobalContext.SetCurrentLocation(arg)
			curr = context.GlobalContext.GetCurrentLocation()
			fmt.Println("You are now in " + curr)

		} else {
			fmt.Println("There is no " + arg + " in the vicinity!")
		}
	}
}
