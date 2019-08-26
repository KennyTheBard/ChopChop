package main

import (
	"bufio"
	"os"

	cli "./cli"
	context "./context"
	handle "./handle"
)

var gContext context.GameContext

func main() {
	gContext.InitContext()
	gContext.GetCommands()["wood"] = handle.TaskFactory("wood", gContext)
	gContext.GetCommands()["pelts"] = handle.TaskFactory("pelts", gContext)
	gContext.GetCommands()["ore"] = handle.TaskFactory("ore", gContext)
	gContext.GetCommands()["inventory"] = handle.InventoryHandler(gContext)

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
