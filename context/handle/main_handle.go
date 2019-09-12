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
		SetPrompt("main")

		if reader.IsInputEqual("exit") {
			break

		} else if reader.IsInputEqual("help") {
			fmt.Println(cli.BuildResponse(
				handleRegister.MapKeys(),
				"What you can do is type:\n * ",
				"\n * ",
				""))

		} else {
			handle := handleRegister.GetHandle(reader.GetInput())

			if handle != nil {
				handle()
			} else {
				fmt.Println("No command known with this name!")
			}
		}
	}
}
