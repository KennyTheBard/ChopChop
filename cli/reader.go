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

func (r *Reader) GetInput() string {
	r.Reload()

	prev := r.Next()
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

func (r *Reader) Next() int {
	r.iter++
	r.Reload()
	return r.iter - 1
}

func (r *Reader) Clear() {
	r.iter = len(r.commands)
}

func (r *Reader) Reload() {
	if r.iter >= len(r.commands) {
		for {
			if r.scanner.Scan() {
				break
			}
		}
		r.commands = strings.Fields(strings.TrimSpace(strings.ToLower(r.scanner.Text())))
		r.iter = 0
	}
}

func (r *Reader) SetPrompt(prompt string) {
	r.prompt = strings.ToUpper(prompt)
}

func (r *Reader) Print(msg string) {
	fmt.Println(msg+" : ", r.commands, r.iter, r.prompt)
}
