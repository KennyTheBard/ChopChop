package component

import (
	"math/rand"
)

type Action struct {
	name       string
	reward     string
	dictionary []string
}

func NewAction(name, reward string, dictionary []string) Action {
	var action Action

	action.name = name
	action.reward = reward
	action.dictionary = dictionary

	return action
}

func (action Action) GetName() string {
	return action.name
}

func (action Action) SetName(name string) {
	action.name = name
}

func (action Action) GetReward() string {
	return action.reward
}

func (action Action) SetReward(reward string) {
	action.reward = reward
}

func (action Action) NextWord() string {
	target := action.dictionary[rand.Intn(len(action.dictionary))]
	return target
}

func (action Action) AddWord(word string) {
	action.dictionary = append(action.dictionary, word)
}
