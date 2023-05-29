package app

import (
	"math/rand"
)

func CreateHypo(battle Battle) *Battle {
	NB := NewBattle()
	for i, b := range battle.Psychics {
		NB.Psychics[i].Hypothesis = 10 + rand.Intn(99-10+1)
		NB.Psychics[i].Reliability = b.Reliability
	}
	return NB
}

func CalculationReliability(battle Battle, number int) *Battle {
	NB := NewBattle()
	for i, b := range battle.Psychics {
		if b.Hypothesis == number {
			NB.Psychics[i].Reliability.Correct = 1 + b.Reliability.Correct
		}
		NB.Psychics[i].Reliability.All = b.Reliability.All + 1
		NB.Psychics[i].Reliability.Percent = (NB.Psychics[i].Reliability.Correct / NB.Psychics[i].Reliability.All) * 100
		NB.Psychics[i].Hypothesis = b.Hypothesis
	}
	return NB
}
