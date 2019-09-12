package main

import (
	"fmt"
	"math/rand"
	"time"

	context "./context"
	component "./context/component"
	handle "./context/handle"
)

func main() {
	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")
	fmt.Println("  Welcome to Woodshack, a text-based game about making the best of what you have!")
	fmt.Println("    In order to learn anything about the current envirment, type 'help'")
	fmt.Println("  At any point there should be a message designed to help you to better understand")
	fmt.Println("    the situation you are in and what can be done. This being said, have fun!")
	fmt.Println(" ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ")

	rand.Seed(time.Now().UTC().UnixNano())

	context.GlobalContext.InitContext()
	actionRegister := context.GlobalContext.GetActionRegister()
	handleRegister := context.GlobalContext.GetHandleRegister()
	itemRegister := context.GlobalContext.GetItemRegister()
	bpRegister := context.GlobalContext.GetBlueprintRegister()
	worldMap := context.GlobalContext.GetWorldMap()

	actionRegister.AddAction(component.NewAction("chop", "wood", 0.6, "axe", []string{"swing", "axe", "saw"}))
	actionRegister.AddAction(component.NewAction("hunt", "meat", 0.4, "spear", []string{"track", "stalk", "shoot"}))
	actionRegister.AddAction(component.NewAction("fish", "fish", 0.25, "", []string{"bait", "throw", "pull"}))
	actionRegister.AddAction(component.NewAction("mine", "ore", 0.15, "pickaxe", []string{"break", "drill", "dig", "polish"}))

	bpRegister.AddBlueprint(component.NewBlueprint("ingot", []string{"ore"}, []int{3}))
	bpRegister.AddBlueprint(component.NewBlueprint("axe", []string{"wood", "ingot"}, []int{2, 2}))
	bpRegister.AddBlueprint(component.NewBlueprint("pickaxe", []string{"wood", "ingot"}, []int{2, 3}))
	bpRegister.AddBlueprint(component.NewBlueprint("spear", []string{"wood", "ingot"}, []int{3, 1}))

	handleRegister.AddHandle("go", handle.TravelHandle)
	handleRegister.AddHandle("do", handle.ActionHandle)
	handleRegister.AddHandle("inv", handle.InventoryHandle)
	handleRegister.AddHandle("craft", handle.CraftHandle)
	handleRegister.AddHandle("help", handle.HelpHandle)
	handleRegister.AddHandle("trade", handle.TradeHandle)

	itemRegister.AddItem(component.NewItem("ingot", 5))

	worldMap.AddLocation(component.NewLocation("forrest", []string{"cave", "river"}, []string{"chop", "hunt"}))
	worldMap.AddLocation(component.NewLocation("river", []string{"forrest"}, []string{"fish"}))
	worldMap.AddLocation(component.NewLocation("cave", []string{"forrest"}, []string{"mine"}))

	context.GlobalContext.SetCurrentLocation("forrest")

	// for testing purposes
	context.GlobalContext.GetInventory().AddItems("wood", 10)
	context.GlobalContext.GetInventory().AddItems("ore", 10)
	context.GlobalContext.GetInventory().AddItems("ingot", 10)

	handle.MainHandle()
}
