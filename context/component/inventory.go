package component

import "strconv"

type Inventory map[string]int

func (inventory Inventory) AddItems(item string, num int) {
	if _, ok := inventory[item]; ok {
		inventory[item] += num
	} else {
		inventory[item] = num
	}
}

func (inventory Inventory) TakeItems(item string, num int) {
	if units, ok := inventory[item]; ok && units >= num {
		inventory[item] -= num
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
	return inventory.HasItemAtLeast(item, 1)
}

func (inventory Inventory) HasItemAtLeast(item string, min int) bool {
	units, ok := inventory[item]
	return ok && units >= min
}

func (inventory Inventory) GetAllItems() []string {
	arr := make([]string, len(inventory))
	counter := 0

	for key, value := range inventory {
		arr[counter] = " " + key + ": " + strconv.Itoa(value)
		counter++
	}

	return arr
}
