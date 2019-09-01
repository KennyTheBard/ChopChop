package handle

import (
	context ".."
)

func MainHandle() {
	reader := context.GlobalContext.GetReader()
	handleRegister := context.GlobalContext.GetHandleRegister()

	for {
		if reader.IsInputEqual("exit") {
			break

		} else {
			handleRegister.GetHandle(reader.GetInput())()
		}
	}
}
