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

	for i, arg := range args {
		builder.Grow(len(arg))
		builder.WriteString(arg)

		if i < len(args)-1 {
			builder.Grow(len(separator))
			builder.WriteString(separator)
		}
	}

	builder.Grow(len(post))
	builder.WriteString(post)

	return builder.String()
}
