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

func HandlerFactory(task string) func(*bufio.Scanner) {
	return func(scanner *bufio.Scanner) {
		target := gContext.GetDictionary().NextWord(task)

		cli.Prompt()
		for scanner.Scan() {
			if strings.EqualFold(target, cli.GetInput(scanner)) {
				gContext.GetInventory().AddItem(task)
				fmt.Println("+1 " + task)
			} else if strings.EqualFold("stop", cli.GetInput(scanner)) {
				break
			}

			target = gContext.GetDictionary().NextWord(task)
			cli.Prompt()
		}
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
	gContext.GetCommands()["wood"] = HandlerFactory("wood")
	gContext.GetCommands()["pelts"] = HandlerFactory("pelts")
	gContext.GetCommands()["ore"] = HandlerFactory("ore")
	gContext.GetCommands()["inventory"] = inventoryHandler

	gContext.GetDictionary()["wood"] = []string{"chop", "swing"}
	gContext.GetDictionary()["pelts"] = []string{"track", "shoot"}
	gContext.GetDictionary()["ore"] = []string{"dig", "smash", "drill"}

	scanner := bufio.NewScanner(os.Stdin)

	cli.Prompt()
	for scanner.Scan() {
		gContext.GetCommands().Interpret(scanner)
		cli.Prompt()
	}

}
