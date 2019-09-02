package handle

import (
	context ".."
)

func MainHandle() {
	handleRegister := context.GlobalContext.GetHandleRegister()
	reader := context.GlobalContext.GetReader()
	reader.SetPrompt("main")

	for {
		if reader.IsInputEqual("exit") {
			break

		} else {
			handleRegister.GetHandle(reader.GetInput())()
		}
	}
}
