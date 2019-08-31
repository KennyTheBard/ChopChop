package component

type WorldMap map[string]Location

func (worldMap WorldMap) AddLocation(locationName string, location Location) {
	worldMap[locationName] = location
}

func (worldMap WorldMap) GetLocation(locationName string) Location {
	return worldMap[locationName]
}

func (worldMap WorldMap) HasLocation(locationName string) bool {
	_, ok := worldMap[locationName]
	return ok
}
