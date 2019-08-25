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

const stdPrompt = ">> "

var prompt string

func interpret(scanner *bufio.Scanner) {
	if handler, ok := gCommands[getInput(scanner)]; ok {
		handler(scanner)
	} else {
		fmt.Println("ERROR")
	}
}

func nextWord(word string) string {
	if words, ok := gDictionary[word]; ok {
		return words[rand.Intn(len(words))]
	}
	return ""
}

func getInput(scanner *bufio.Scanner) string {
	return strings.TrimSpace(strings.ToLower(scanner.Text()))
}

func requestInput(target string) {
	fmt.Println("Type '" + target + "'")
	fmt.Print(prompt)
}

func chopWoodHandler(scanner *bufio.Scanner) {
	word := "wood"
	target := nextWord(word)
	requestInput(target)

	for scanner.Scan() {
		if strings.EqualFold(target, getInput(scanner)) {
			gInventory["wood"]++
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
			if item, ok := gInventory[getInput(scanner)]; ok {
				fmt.Println(getInput(scanner) + ": " + strconv.Itoa(item))
			} else {
				fmt.Println("None")
			}
		}

	}

}

type commands map[string]func(*bufio.Scanner)

var gCommands commands

type dictionary map[string][]string

var gDictionary dictionary

type inventory map[string]int

var gInventory inventory

func main() {
	playing = true
	prompt = stdPrompt

	gCommands = make(commands)
	gCommands["wood"] = chopWoodHandler
	gCommands["inventory"] = inventoryHandler

	gDictionary = make(dictionary)
	gDictionary["wood"] = []string{"chop", "swing"}

	gInventory = make(inventory)
	gInventory["wood"] = 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(prompt)
	for scanner.Scan() && playing {
		interpret(scanner)
		fmt.Print(prompt)
	}

}
