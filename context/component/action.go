package component

import (
	"math/rand"
)

type Action struct {
	name, reward string
	chance       float32
	tool         string
	success      bool
	dictionary   []string
}

func NewAction(name, reward string, chance float32, tool string, dictionary []string) Action {
	var action Action

	action.name = name
	action.reward = reward
	action.chance = chance
	action.tool = tool
	action.success = false
	action.dictionary = dictionary

	return action
}

func (action Action) GetName() string {
	return action.name
}

func (action *Action) SetName(name string) {
	action.name = name
}

func (action Action) GetReward() string {
	return action.reward
}

func (action *Action) SetReward(reward string) {
	action.reward = reward
}

func (action Action) GetTool() string {
	return action.tool
}

func (action *Action) SetTool(tool string) {
	action.tool = tool
}

func (action Action) IsSuccessful() bool {
	return action.success
}

func (action *Action) NextWord() string {
	if rand.Float32() < action.chance {
		action.success = true
	} else {
		action.success = false
	}
	target := action.dictionary[rand.Intn(len(action.dictionary))]
	return target
}

func (action Action) AddWord(word string) {
	action.dictionary = append(action.dictionary, word)
}
