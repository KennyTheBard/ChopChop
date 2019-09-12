package handle

import (
	context ".."
	cli "../../cli"
)

func SetPrompt(prompt string) {
	cli.SetPrompt(prompt, context.GlobalContext.GetCurrentLocation())
}
