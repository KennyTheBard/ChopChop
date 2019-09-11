package component

type Handle func()

type HandleRegister map[string]Handle

func (hRegister HandleRegister) GetHandle(handle string) Handle {
	return hRegister[handle]
}

func (hRegister HandleRegister) HasHandle(handle string) bool {
	_, ok := hRegister[handle]
	return ok
}

func (hRegister HandleRegister) AddHandle(handleName string, handle Handle) {
	hRegister[handleName] = handle
}

func (hRegister HandleRegister) MapKeys() []string {
	keys := make([]string, 0, len(hRegister))
	for k := range hRegister {
		keys = append(keys, k)
	}
	return keys
}
