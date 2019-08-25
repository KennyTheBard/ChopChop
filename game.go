package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	cli "./cli"
	context "./context"
)

func chopWoodHandler(scanner *bufio.Scanner) {
	word := "wood"
	target := gContext.GetDictionary().NextWord(word)

	for scanner.Scan() {
		if strings.EqualFold(target, cli.GetInput(scanner)) {
			gContext.GetInventory().AddItem("wood")
			fmt.Println("Chopped 1 Wood!")
		} else if strings.EqualFold("stop", cli.GetInput(scanner)) {
			break
		}
		target = gContext.GetDictionary().NextWord(word)
	}
}

func inventoryHandler(scanner *bufio.Scanner) {
	fmt.Print("What are you looking for in your inventory?\n>> ")

	for scanner.Scan() {
		if strings.EqualFold("stop", cli.GetInput(scanner)) {
			break
		} else {
			if item, ok := gContext.GetInventory()[cli.GetInput(scanner)]; ok {
				fmt.Println(cli.GetInput(scanner) + ": " + strconv.Itoa(item))
			} else {
				fmt.Println("None")
			}
		}
		cli.Prompt()
	}

}

var gContext context.GameContext

func main() {

	gContext.InitContext()
	gContext.GetCommands()["wood"] = chopWoodHandler
	gContext.GetCommands()["inventory"] = inventoryHandler
	gContext.GetDictionary()["wood"] = []string{"chop", "swing"}

	scanner := bufio.NewScanner(os.Stdin)

	cli.Prompt()
	for scanner.Scan() {
		gContext.GetCommands().Interpret(scanner)
		cli.Prompt()
	}

}
