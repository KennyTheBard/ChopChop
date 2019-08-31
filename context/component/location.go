package component

type Location struct {
	name              string
	adjacentLocations []string
	availableActions  []string
}

func NewLocation(name string, adjacentLocations, availableActions []string) Location {
	var location Location

	location.name = name
	location.adjacentLocations = adjacentLocations
	location.availableActions = availableActions

	return location
}

func (location Location) GetName() string {
	return location.name
}

func (location Location) SetName(name string) {
	location.name = name
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

func (location Location) HasAvailableAction(action string) bool {
	for _, availableAction := range location.availableActions {
		if action == availableAction {
			return true
		}
	}

	return false
}
