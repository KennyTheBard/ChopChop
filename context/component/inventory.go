package component

type Inventory map[string]int

func (inventory Inventory) AddItem(item string) {
	if _, ok := inventory[item]; ok {
		inventory[item]++
	} else {
		inventory[item] = 1
	}
}

func (inventory Inventory) GetItem(item string) int {
	if units, ok := inventory[item]; ok {
		return units
	} else {
		return 0
	}
}

func (inventory Inventory) HasItem(item string) bool {
	units, ok := inventory[item]
	return ok && units > 0
}
