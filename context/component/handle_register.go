package component

type Handle func()

type HandleRegister struct {
	register map[string]Handle
	fallback string
}

func (hRegister *HandleRegister) Init(fallbackHandleName string) {
	hRegister.register = make(map[string]Handle)
	hRegister.fallback = fallbackHandleName
}

func (hRegister HandleRegister) GetHandle(handle string) Handle {
	if handle, ok := hRegister.register[handle]; ok {
		return handle
	} else {
		return hRegister.register[hRegister.fallback]
	}
}

func (hRegister HandleRegister) HasHandle(handle string) bool {
	_, ok := hRegister.register[handle]
	return ok
}

func (hRegister HandleRegister) AddHandle(handleName string, handle Handle) {
	hRegister.register[handleName] = handle
}
