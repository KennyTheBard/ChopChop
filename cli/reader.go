package cli

import (
	"bufio"
	"os"
	"strings"
)

type Reader struct {
	scanner  *bufio.Scanner
	commands []string
	iter     int
}

func (r *Reader) Init() {
	r.scanner = bufio.NewScanner(os.Stdin)
	r.commands = []string{}
	r.iter = 0
}

func (r Reader) GetInput() string {
	r.Reload()

	prev := r.Next()
	return r.commands[prev]
}

func (r Reader) IsInputEqual(cmd string) bool {
	r.Reload()

	return r.commands[r.iter] == strings.ToLower(cmd)
}

func (r *Reader) Next() int {
	r.iter++
	return r.iter - 1
}

func (r *Reader) Reload() {
	if r.iter >= len(r.commands) {
		Prompt()
		r.scanner.Scan()
		r.commands = strings.Fields(strings.TrimSpace(strings.ToLower(r.scanner.Text())))
		r.iter = 0
	}
}
