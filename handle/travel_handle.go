package handle

import (
	"fmt"
	"strings"

	cli "../cli"
	context "../context"
)

func TravelHandle(args []string) {
	for _, arg := range args {
		curr := context.GlobalContext.GetCurrentLocation()
		wMap := context.GlobalContext.GetWorldMap()

		if strings.EqualFold(arg, "where") {
			fmt.Println(cli.BuildResponse(
				wMap.GetLocation(curr).GetAdjacentLocations(),
				" * ",
				"\n * ",
				""))

		} else {
			if wMap.GetLocation(curr).HasAdjacentLocation(arg) {
				context.GlobalContext.SetCurrentLocation(arg)
				fmt.Println("You are now in " + arg)

			} else {
				fmt.Println("ERROR [" + arg + "]: No such location in vicinity!")
				break
			}
		}
	}
}
