package handle

import (
	context ".."
	cli "../../cli"
)

func MainHandle() {
	handleRegister := context.GlobalContext.GetHandleRegister()
	reader := context.GlobalContext.GetReader()
	cli.SetPrompt("main")

	for {
		if reader.IsInputEqual("exit") {
			break

		} else {
			handleRegister.GetHandle(reader.GetInput())()
		}
	}
}
