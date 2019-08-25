package context

type Inventory map[string]int

func (inventory Inventory) AddItem(item string) {
	if _, ok := inventory[item]; ok {
		inventory[item]++
	} else {
		inventory[item] = 1
	}
}
