package component

type Blueprint struct {
	item string
	cost map[string]int
}

func NewBlueprint(item string, items []string, costs []int) Blueprint {
	var bp Blueprint

	bp.item = item
	bp.cost = make(map[string]int)
	for i := 0; i < len(items) && i < len(costs); i++ {
		bp.cost[items[i]] = costs[i]
	}

	return bp
}

func (bp Blueprint) GetItem() string {
	return bp.item
}

func (bp *Blueprint) SetItem(item string) {
	bp.item = item
}

func (bp Blueprint) SetCost(item string, value int) {
	bp.cost[item] = value
}

func (bp Blueprint) GetCost(item string) int {
	return bp.cost[item]
}

func (bp Blueprint) GetAllCosts() map[string]int {
	return bp.cost
}
