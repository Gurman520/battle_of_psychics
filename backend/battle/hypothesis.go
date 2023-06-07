package battle

import (
	"math/rand"
)

// The function creates hypotheses
func CreateHypotheses(battle Battle) *Battle {
	NB := NewBattle()
	for i, b := range battle.Psychics {
		NB.Psychics[i].Hypothesis = append(b.Hypothesis, 10+rand.Intn(99-10+1))
		NB.Psychics[i].Reliability = b.Reliability
	}
	NB.UserNumber = battle.UserNumber
	return NB
}
