package handle

import (
	"fmt"
	"strconv"

	context ".."
	cli "../../cli"
)

func TradeHandle() {
	inventory := context.GlobalContext.GetInventory()
	reg := context.GlobalContext.GetItemRegister()
	reader := context.GlobalContext.GetReader()

	for {
		SetPrompt("trade")

		if reader.IsInputEqual("help") {
			items := make([]string, len(inventory))

			i := 0
			for item := range inventory {
				items[i] = item + " -> "
				if reg.HasItem(item) {
					items[i] += strconv.Itoa(reg.GetItem(item).GetValue())
				} else {
					items[i] += "0"
				}

				i++
			}

			fmt.Println(cli.BuildResponse(
				items,
				"Trade prices ares:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else {
			fmt.Println("There is no " + reader.GetInput() + " in the vicinity!")
		}
	}
}
