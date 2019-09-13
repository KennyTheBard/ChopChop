package handle

import (
	"fmt"
	"strconv"

	context ".."
	cli "../../cli"
	npc "../npc"
)

func TradeHandle() {
	reader := context.GlobalContext.GetReader()
	merchant := context.GlobalContext.GetMerchant()

	for {
		SetPrompt("trade")

		if reader.IsInputEqual("help") {
			items := make([]string, len(merchant.GetAllOffers()))

			i := 0
			for item, offer := range merchant.GetAllOffers() {
				items[i] = strconv.Itoa(offer) + " x " + item + " for 1 " + merchant.GetItem()
				i++
			}

			fmt.Println(cli.BuildResponse(
				items,
				"Trade prices are:\n * ",
				"\n * ",
				""))

		} else if reader.IsInputEqual("back") {
			break

		} else if arg := reader.GetInput(); merchant.HasOffer(arg) {
			if tradeItem(merchant, arg) {
				fmt.Println("You traded " + strconv.Itoa(merchant.GetOffer(arg)) + " x " + arg + " for 1 " + merchant.GetItem())
			} else {
				fmt.Println("You don't have enough " + arg + " to trade!")
			}

		} else {
			fmt.Println("This merchant does not trade for " + arg + "!")
		}
	}
}

func tradeItem(merchant npc.Merchant, item string) bool {
	inventory := context.GlobalContext.GetInventory()

	if !inventory.HasItemAtLeast(item, merchant.GetOffer(item)) {
		return false
	}

	inventory.TakeItems(item, merchant.GetOffer(item))
	inventory.AddItems(merchant.GetItem(), 1)
	return true
}
