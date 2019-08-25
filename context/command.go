package context

import (
	"bufio"
	"fmt"

	cli "../cli"
)

type Commands map[string]func(*bufio.Scanner)

func (commands Commands) Interpret(scanner *bufio.Scanner) {
	if handler, ok := commands[cli.GetInput(scanner)]; ok {
		handler(scanner)
	} else {
		fmt.Println("ERROR")
	}
}
