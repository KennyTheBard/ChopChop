package main

import (
	context "./context"
	component "./context/component"
	handle "./context/handle"
)

func main() {

	context.GlobalContext.InitContext()
	actionRegister := context.GlobalContext.GetActionRegister()
	handleRegister := context.GlobalContext.GetHandleRegister()
	bpRegister := context.GlobalContext.GetBlueprintRegister()
	worldMap := context.GlobalContext.GetWorldMap()

	actionRegister.AddAction(component.NewAction("chop", "wood", []string{"swing", "axe", "saw"}))
	actionRegister.AddAction(component.NewAction("hunt", "meat", []string{"track", "stalk", "shoot"}))
	actionRegister.AddAction(component.NewAction("fish", "fish", []string{"bait", "throw", "pull"}))
	actionRegister.AddAction(component.NewAction("mine", "ore", []string{"break", "drill", "dig", "polish"}))

	bpRegister.AddBlueprint(component.NewBlueprint("ingot", []string{"ore"}, []int{3}))
	bpRegister.AddBlueprint(component.NewBlueprint("axe", []string{"wood", "ingot"}, []int{2, 2}))
	bpRegister.AddBlueprint(component.NewBlueprint("pickaxe", []string{"wood", "ingot"}, []int{2, 3}))
	bpRegister.AddBlueprint(component.NewBlueprint("spear", []string{"wood", "ingot"}, []int{3, 1}))

	handleRegister.Init("help")
	handleRegister.AddHandle("go", handle.TravelHandle)
	handleRegister.AddHandle("do", handle.ActionHandle)
	handleRegister.AddHandle("inv", handle.InventoryHandle)
	handleRegister.AddHandle("craft", handle.WorkbenchHandle)
	handleRegister.AddHandle("help", handle.HelpHandle)
	context.GlobalContext.SetHandleRegister(handleRegister)

	worldMap.AddLocation(component.NewLocation("forrest", []string{"cave", "river"}, []string{"chop", "hunt"}))
	worldMap.AddLocation(component.NewLocation("river", []string{"forrest"}, []string{"fish"}))
	worldMap.AddLocation(component.NewLocation("cave", []string{"forrest"}, []string{"mine"}))

	context.GlobalContext.SetCurrentLocation("forrest")

	handle.MainHandle()
}
