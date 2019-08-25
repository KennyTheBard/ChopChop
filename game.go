package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var playing bool

func Prompt() {
	fmt.Print(">> ")
}

func interpret(scanner *bufio.Scanner) {
	if handler, ok := context.gCommands[getInput(scanner)]; ok {
		handler(scanner)
	} else {
		fmt.Println("ERROR")
	}
}

func nextWord(word string) string {
	if words, ok := context.gDictionary[word]; ok {
		return words[rand.Intn(len(words))]
	}
	return ""
}

func getInput(scanner *bufio.Scanner) string {
	return strings.TrimSpace(strings.ToLower(scanner.Text()))
}

func requestInput(target string) {
	fmt.Println("Type '" + target + "'")
	Prompt()
}

func addItem(item string) {
	if _, ok := context.gInventory[item]; ok {
		context.gInventory[item]++
	} else {
		context.gInventory[item] = 1
	}
}

func chopWoodHandler(scanner *bufio.Scanner) {
	word := "wood"
	target := nextWord(word)
	requestInput(target)

	for scanner.Scan() {
		if strings.EqualFold(target, getInput(scanner)) {
			addItem("wood")
			fmt.Println("Chopped 1 Wood!")
		} else if strings.EqualFold("stop", getInput(scanner)) {
			break
		}
		target = nextWord(word)
		requestInput(target)
	}
}

func inventoryHandler(scanner *bufio.Scanner) {
	fmt.Print("What are you looking for in your inventory?\n>> ")

	for scanner.Scan() {
		if strings.EqualFold("stop", getInput(scanner)) {
			break
		} else {
			if item, ok := context.gInventory[getInput(scanner)]; ok {
				fmt.Println(getInput(scanner) + ": " + strconv.Itoa(item))
			} else {
				fmt.Println("None")
			}
		}
		Prompt()
	}

}

type Commands map[string]func(*bufio.Scanner)

type Dictionary map[string][]string

type Inventory map[string]int

type GameContext struct {
	gCommands   Commands
	gDictionary Dictionary
	gInventory  Inventory
}

var context GameContext

func (context *GameContext) InitContext() {
	context.gCommands = make(Commands)
	context.gDictionary = make(Dictionary)
	context.gInventory = make(Inventory)
}

func main() {
	playing = true

	context.InitContext()
	context.gCommands["wood"] = chopWoodHandler
	context.gCommands["inventory"] = inventoryHandler
	context.gDictionary["wood"] = []string{"chop", "swing"}

	scanner := bufio.NewScanner(os.Stdin)

	Prompt()
	for scanner.Scan() && playing {
		interpret(scanner)
		Prompt()
	}

}
