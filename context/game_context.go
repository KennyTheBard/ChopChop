package context

import (
	component "./component"
)

type GameContext struct {
	inventory       component.Inventory
	worldMap        component.WorldMap
	actionRegister  component.ActionRegister
	currentLocation string
}

var GlobalContext GameContext

func (context *GameContext) InitContext() {
	context.inventory = make(component.Inventory)
	context.worldMap = make(component.WorldMap)
	context.actionRegister = make(component.ActionRegister)
}

func (context GameContext) GetInventory() component.Inventory {
	return context.inventory
}

func (context GameContext) GetWorldMap() component.WorldMap {
	return context.worldMap
}

func (context GameContext) GetActionRegister() component.ActionRegister {
	return context.actionRegister
}

func (context GameContext) GetCurrentLocation() string {
	return context.currentLocation
}

func (context GameContext) SetCurrentLocation(location string) {
	context.currentLocation = location
}
