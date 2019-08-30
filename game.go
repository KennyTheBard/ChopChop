package main

import (
	"bufio"
	"os"

	cli "./cli"
	handle "./handle"
)

func main() {

	handle.GlobalContext.InitContext()
	handle.GlobalContext.GetCommands()["wood"] = handle.TaskFactory("wood")
	handle.GlobalContext.GetCommands()["pelts"] = handle.TaskFactory("pelts")
	handle.GlobalContext.GetCommands()["ore"] = handle.TaskFactory("ore")
	handle.GlobalContext.GetCommands()["inventory"] = handle.InventoryHandler()

	handle.GlobalContext.GetDictionary()["wood"] = []string{"chop", "swing"}
	handle.GlobalContext.GetDictionary()["pelts"] = []string{"track", "shoot"}
	handle.GlobalContext.GetDictionary()["ore"] = []string{"dig", "smash", "drill"}

	scanner := bufio.NewScanner(os.Stdin)

	cli.Prompt()
	for scanner.Scan() {
		handle.GlobalContext.GetCommands().Interpret(scanner)
		cli.Prompt()
	}

}
