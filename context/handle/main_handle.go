package handle

import (
	"fmt"

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

		} else if reader.IsInputEqual("help") {
			fmt.Println(cli.BuildResponse(
				handleRegister.MapKeys(),
				"What you can do is type:\n * ",
				"\n * ",
				""))

		} else {
			handleRegister.GetHandle(reader.GetInput())()
		}
	}
}
