package component

type Location struct {
	adjacentLocations []string
	availableActions  []string
}

func (location Location) GetAdjacentLocations() []string {
	return location.adjacentLocations
}

func (location Location) AddAdjacentLocation(loc string) {
	location.adjacentLocations = append(location.adjacentLocations, loc)
}

func (location Location) HasAdjacentLocation(loc string) bool {
	for _, adjLocation := range location.adjacentLocations {
		if loc == adjLocation {
			return true
		}
	}

	return false
}

func (location Location) GetAvailableActions() []string {
	return location.availableActions
}

func (location Location) AddAvailableAction(action string) {
	location.availableActions = append(location.availableActions, action)
}
