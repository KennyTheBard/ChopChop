package handle

import (
	context ".."
	cli "../../cli"
)

func MainHandle() {
	handleRegister := context.GlobalContext.GetHandleRegister()
	reader := context.GlobalContext.GetReader()

	for {
		cli.SetPrompt("main")

		if reader.IsInputEqual("exit") {
			break

		} else {
			handleRegister.GetHandle(reader.GetInput())()
		}
	}
}
