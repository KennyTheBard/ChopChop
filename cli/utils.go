package cli

import (
	"bufio"
	"fmt"
	"strings"
)

func Prompt() {
	fmt.Print(">> ")
}

func GetInput(scanner *bufio.Scanner) string {
	return strings.TrimSpace(strings.ToLower(scanner.Text()))
}
