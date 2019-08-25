package context

type GameContext struct {
	gCommands   Commands
	gDictionary Dictionary
	gInventory  Inventory
}

func (context *GameContext) InitContext() {
	context.gCommands = make(Commands)
	context.gDictionary = make(Dictionary)
	context.gInventory = make(Inventory)
}

func (context GameContext) GetCommands() Commands {
	return context.gCommands
}

func (context GameContext) GetDictionary() Dictionary {
	return context.gDictionary
}

func (context GameContext) GetInventory() Inventory {
	return context.gInventory
}
