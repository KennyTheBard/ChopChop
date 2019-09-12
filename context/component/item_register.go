package component

type ItemRegister map[string]Item

func (reg ItemRegister) GetItem(name string) Item {
	return reg[name]
}

func (reg ItemRegister) HasItem(name string) bool {
	_, ok := reg[name]
	return ok
}

func (reg ItemRegister) AddItem(item Item) {
	reg[item.GetName()] = item
}
