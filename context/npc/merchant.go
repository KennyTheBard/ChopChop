package npc

type Merchant struct {
	item   string
	offers map[string]int
}

func NewMerchant(item string, offers map[string]int) Merchant {
	var m Merchant

	m.item = item
	m.offers = offers

	return m
}

func (m Merchant) GetItem() string {
	return m.item
}

func (m *Merchant) SetItem(item string) {
	m.item = item
}

func (m Merchant) GetOffer(item string) int {
	return m.offers[item]
}

func (m Merchant) HasOffer(item string) bool {
	_, ok := m.offers[item]
	return ok
}

func (m *Merchant) AddOffer(item string, offer int) {
	m.offers[item] = offer
}

func (m Merchant) GetAllOffers() map[string]int {
	return m.offers
}
