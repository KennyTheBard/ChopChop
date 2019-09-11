package context

import (
	"fmt"

	cli "../cli"
	component "./component"
)

type GameContext struct {
	inventory         component.Inventory
	actionRegister    component.ActionRegister
	handleRegister    component.HandleRegister
	itemRegister      component.ItemRegister
	blueprintRegister component.BlueprintRegister
	worldMap          component.WorldMap
	reader            *cli.Reader
	currentLocation   string
}

var GlobalContext GameContext

func (context *GameContext) InitContext() {
	context.inventory = make(component.Inventory)
	context.actionRegister = make(component.ActionRegister)
	context.handleRegister = make(component.HandleRegister)
	context.itemRegister = make(component.ItemRegister)
	context.blueprintRegister = make(component.BlueprintRegister)
	context.worldMap = make(component.WorldMap)
	context.reader = new(cli.Reader)

	context.reader.Init()
}

func (context GameContext) GetInventory() component.Inventory {
	return context.inventory
}

func (context GameContext) GetActionRegister() component.ActionRegister {
	return context.actionRegister
}

func (context GameContext) GetHandleRegister() component.HandleRegister {
	return context.handleRegister
}

func (context GameContext) GetItemRegister() component.ItemRegister {
	return context.itemRegister
}

func (context GameContext) GetBlueprintRegister() component.BlueprintRegister {
	return context.blueprintRegister
}

func (context GameContext) GetWorldMap() component.WorldMap {
	return context.worldMap
}

func (context GameContext) GetReader() *cli.Reader {
	return context.reader
}

func (context GameContext) GetCurrentLocation() string {
	return context.currentLocation
}

func (context *GameContext) SetCurrentLocation(location string) {
	context.currentLocation = location
}

func (context GameContext) Print() {
	fmt.Println(context.inventory)
	fmt.Println(context.actionRegister)
	fmt.Println(context.worldMap)
	fmt.Println(context.currentLocation)
}
