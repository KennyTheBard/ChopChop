package component

type WorldMap map[string]Location

func (worldMap WorldMap) AddLocation(location Location) {
	worldMap[location.GetName()] = location
}

func (worldMap WorldMap) GetLocation(locationName string) Location {
	return worldMap[locationName]
}

func (worldMap WorldMap) HasLocation(locationName string) bool {
	_, ok := worldMap[locationName]
	return ok
}
