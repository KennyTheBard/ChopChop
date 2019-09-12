package npc

type Merchant struct {
	item  string
	price int
}

func NewMerchant(item string, price int) Merchant {
	var m Merchant

	m.item = item
	m.price = price

	return m
}

func (m Merchant) GetItem() string {
	return m.item
}

func (m Merchant) GetPrice() int {
	return m.price
}
