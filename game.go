package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var playing bool

const stdPrompt = ">> "

var prompt string

func interpret(cmd string, cmds commands, inv inventory) string {
	if strings.EqualFold(cmds.chopCmd, cmd) {
		inv.wood++
		return "1 Wood chopped"
	}
	return "ERROR"
}

type commands struct {
	chopCmd string
}

type inventory struct {
	wood int
}

func main() {
	playing = true
	prompt = stdPrompt

	cmds := commands{}
	cmds.chopCmd = "chop"

	inv := inventory{}
	inv.wood = 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	for scanner.Scan() && playing {
		fmt.Println(interpret(scanner.Text(), cmds, inv))
		fmt.Print(prompt)
	}

}
