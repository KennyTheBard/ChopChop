package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var playing bool

const stdPrompt = ">> "

var prompt string

func interpret(cmd string, cmds commands, inv *inventory) string {
	if strings.EqualFold(cmds.chopCmd, cmd) {
		inv.wood++
		return "1 Wood chopped"
	} else if strings.EqualFold(cmds.invCmd, cmd) {
		return "There is " + strconv.Itoa(inv.wood) + " Wood in the bag"
	}
	return "ERROR"
}

type commands struct {
	chopCmd string
	invCmd  string
}

type inventory struct {
	wood int
}

func main() {
	playing = true
	prompt = stdPrompt

	cmds := commands{}
	cmds.chopCmd = "chop"
	cmds.invCmd = "eval"

	inv := inventory{}
	inv.wood = 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	for scanner.Scan() && playing {
		fmt.Println(interpret(scanner.Text(), cmds, &inv))
		fmt.Print(prompt)
	}

}
