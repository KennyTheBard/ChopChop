package context

import (
	"fmt"
	"math/rand"
)

type Dictionary map[string][]string

func (dictionary Dictionary) NextWord(word string) string {
	if words, ok := dictionary[word]; ok {
		target := words[rand.Intn(len(words))]
		fmt.Println("Type '" + target + "'")
		return target
	}

	fmt.Println("No dictionary found for the given word")
	return ""
}
