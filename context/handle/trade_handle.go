package handle

import (
	"fmt"
	"math"
	"strconv"

	context ".."
	cli "../../cli"
)

func TradeHandle() {
	inventory := context.GlobalContext.GetInventory()
	reg := context.GlobalContext.GetItemRegister()
	reader := context.GlobalContext.GetReader()
	merchant := context.GlobalContext.GetMerchant()

	for {
		SetPrompt("trade")

		if reader.IsInputEqual("help") {
			items := make([]string, len(inventory))

			i := 0
			for item := range inventory {
				if reg.HasItem(item) {
					items[i] = strconv.Itoa(int(math.Ceil(float64(merchant.GetPrice())/float64(reg.GetItem(item).GetValue())))) + "x " + item + " for 1 " + merchant.GetItem()
				} else {
					items[i] = "Can't be bought with " + item
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
