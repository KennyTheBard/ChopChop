package component

type Item struct {
	name  string
	value int
}

func NewItem(name string, value int) Item {
	var item Item

	item.name = name
	item.value = value

	return item
}

func (i Item) GetName() string {
	return i.name
}

func (i *Item) SetName(name string) {
	i.name = name
}

func (i Item) GetValue() int {
	return i.value
}

func (i *Item) SetValue(value int) {
	i.value = value
}
