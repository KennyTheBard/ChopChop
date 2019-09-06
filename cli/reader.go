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

func (r *Reader) GetInput() string {
	r.Reload()

	prev := r.iter
	r.iter++
	return r.commands[prev]
}

func (r *Reader) IsInputEqual(cmd string) bool {
	r.Reload()

	if r.commands[r.iter] == strings.ToLower(cmd) {
		r.iter++
		return true
	} else {
		return false
	}
}

func (r *Reader) Clear() {
	r.iter = len(r.commands)
}

func (r *Reader) Reload() {
	if r.iter >= len(r.commands) {
		Prompt()
		for {
			if r.scanner.Scan() {
				break
			}
		}
		r.commands = strings.Fields(strings.TrimSpace(strings.ToLower(r.scanner.Text())))
		r.iter = 0
	}
}
