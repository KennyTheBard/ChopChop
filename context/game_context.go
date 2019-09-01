package context

import (
	"bufio"
	"fmt"
	"os"

	component "./component"
)

type GameContext struct {
	inventory       component.Inventory
	actionRegister  component.ActionRegister
	handleRegister  component.HandleRegister
	worldMap        component.WorldMap
	scanner         *bufio.Scanner
	currentLocation string
}

var GlobalContext GameContext

func (context *GameContext) InitContext() {
	context.inventory = make(component.Inventory)
	context.actionRegister = make(component.ActionRegister)
	context.worldMap = make(component.WorldMap)

	context.scanner = bufio.NewScanner(os.Stdin)
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

func (context GameContext) GetWorldMap() component.WorldMap {
	return context.worldMap
}

func (context GameContext) GetScanner() *bufio.Scanner {
	return context.scanner
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
