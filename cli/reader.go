package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader struct {
	scanner  *bufio.Scanner
	commands []string
	iter     int
	prompt   string
}

func (r *Reader) Init() {
	r.scanner = bufio.NewScanner(os.Stdin)
	r.commands = []string{}
	r.iter = 0
	r.prompt = ""
}

func (r Reader) GetInput() string {
	r.Reload()

	prev := r.Next()
	fmt.Println(prev, r.commands, r.iter)
	return r.commands[prev]
}

func (r Reader) IsInputEqual(cmd string) bool {
	r.Reload()

	fmt.Println(cmd, r.commands, r.iter)
	return r.commands[r.iter] == strings.ToLower(cmd)
}

func (r *Reader) Next() int {
	r.iter++
	return r.iter - 1
}

func (r *Reader) Reload() {
	if r.iter >= len(r.commands) {
		Prompt(r.prompt)
		r.scanner.Scan()
		r.commands = strings.Fields(strings.TrimSpace(strings.ToLower(r.scanner.Text())))
		r.iter = 0
	}
}

func (r *Reader) SetPrompt(prompt string) {
	fmt.Println(prompt, r.commands, r.iter)
	r.prompt = strings.ToUpper(prompt)
}
