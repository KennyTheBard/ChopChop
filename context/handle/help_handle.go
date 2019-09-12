package handle

import (
	"fmt"

	context ".."
)

func HelpHandle() {
	SetPrompt("help")
	fmt.Println("This is Woodshack help menu. If you wish to learn more about anything, use it.")
	context.GlobalContext.GetReader().Clear()
}
