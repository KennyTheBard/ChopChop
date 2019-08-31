package component

type ActionRegister map[string]Action

func (register ActionRegister) GetAction(action string) Action {
	return register[action]
}

func (register ActionRegister) HasAction(action string) bool {
	_, ok := register[action]
	return ok
}

func (register ActionRegister) AddAction(action Action) {
	register[action.GetName()] = action
}
