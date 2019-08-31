package component

import (
	"math/rand"
)

type Action struct {
	name       string
	dictionary []string
	reward     string
}

func (action Action) GetName() string {
	return action.name
}

func (action Action) SetName(name string) {
	action.name = name
}

func (action Action) NextWord() string {
	target := action.dictionary[rand.Intn(len(action.dictionary))]
	return target
}

func (action Action) AddWord(word string) {
	action.dictionary = append(action.dictionary, word)
}

func (action Action) GetReward() string {
	return action.reward
}

func (action Action) SetReward(reward string) {
	action.reward = reward
}
