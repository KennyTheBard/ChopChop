package component

type ActionRegister map[string]Action

func (register ActionRegister) GetAction(action string) Action {
	return register[action]
}

func (register ActionRegister) HasAction(action string) bool {
	_, ok := register[action]
	return ok
}

func (register ActionRegister) AddAction(actionName string, action Action) {
	register[actionName] = action
}
