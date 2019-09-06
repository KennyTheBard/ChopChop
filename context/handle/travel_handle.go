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
	cli.SetPrompt("travel")

	for {
		if reader.IsInputEqual("where") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAdjacentLocations(),
				"Close locations are:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else if arg := reader.GetInput(); wMap.GetLocation(curr).HasAdjacentLocation(arg) {
			context.GlobalContext.SetCurrentLocation(arg)
			fmt.Println("You are now in " + arg)

		} else {
			fmt.Println("ERROR [" + arg + "]: No such location in vicinity!")
		}
	}
}
