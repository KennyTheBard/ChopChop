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

func BuildResponse(args []string, pre, separator, post string) string {
	var builder strings.Builder

	builder.Grow(len(pre))
	builder.WriteString(pre)

	for _, arg := range args {
		builder.Grow(len(arg) + len(separator))
		builder.WriteString(arg)
		builder.WriteString(separator)
	}

	builder.Grow(len(post))
	builder.WriteString(post)

	return builder.String()
}
