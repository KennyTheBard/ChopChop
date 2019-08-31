package main

import (
	"bufio"
	"os"

	cli "./cli"
	context "./context"
	handle "./handle"
)

func main() {

	context.GlobalContext.InitContext()
	context.GlobalContext.GetCommands()["wood"] = handle.TaskFactory("wood")
	context.GlobalContext.GetCommands()["pelts"] = handle.TaskFactory("pelts")
	context.GlobalContext.GetCommands()["ore"] = handle.TaskFactory("ore")
	context.GlobalContext.GetCommands()["inventory"] = handle.InventoryHandler()

	context.GlobalContext.GetDictionary()["wood"] = []string{"chop", "swing"}
	context.GlobalContext.GetDictionary()["pelts"] = []string{"track", "shoot"}
	context.GlobalContext.GetDictionary()["ore"] = []string{"dig", "smash", "drill"}

	scanner := bufio.NewScanner(os.Stdin)

	cli.Prompt()
	for scanner.Scan() {
		context.GlobalContext.GetCommands().Interpret(scanner)
		cli.Prompt()
	}

}
