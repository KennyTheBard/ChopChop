package main

import (
	"strings"

	cli "./cli"
	context "./context"
	component "./context/component"
	handle "./context/handle"
)

func main() {

	context.GlobalContext.InitContext()
	actionRegister := context.GlobalContext.GetActionRegister()
	handleRegister := context.GlobalContext.GetHandleRegister()
	worldMap := context.GlobalContext.GetWorldMap()

	actionRegister.AddAction(component.NewAction("chop", "wood", []string{"swing", "axe", "saw"}))
	actionRegister.AddAction(component.NewAction("hunt", "meat", []string{"track", "stalk", "shoot"}))
	actionRegister.AddAction(component.NewAction("fish", "fish", []string{"bait", "throw", "pull"}))
	actionRegister.AddAction(component.NewAction("mine", "ore", []string{"break", "drill", "dig", "polish"}))

	handleRegister.Init("help")
	handleRegister.AddHandle("go", handle.TravelHandle)
	handleRegister.AddHandle("do", handle.ActionHandle)
	handleRegister.AddHandle("inv", handle.InventoryHandle)
	handleRegister.AddHandle("help", handle.HelpHandle)

	worldMap.AddLocation(component.NewLocation("forrest", []string{"cave", "river"}, []string{"chop", "hunt"}))
	worldMap.AddLocation(component.NewLocation("river", []string{"forrest"}, []string{"fish"}))
	worldMap.AddLocation(component.NewLocation("cave", []string{"forrest"}, []string{"mine"}))

	context.GlobalContext.SetCurrentLocation("forrest")

	scanner := context.GlobalContext.GetScanner()
	for {
		cli.Prompt()
		scanner.Scan()

		if cli.GetInput(scanner) == "exit" {
			break

		} else {
			args := strings.Split(cli.GetInput(scanner), " ")
			reminder := args[1:]

			handleRegister.GetHandle(args[0])(reminder)
		}
	}

}
