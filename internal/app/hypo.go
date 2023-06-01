package app

import (
	"math/rand"
)

// Getting new Psychic Predictions
func CreateHypo(battle Battle) *Battle {
	NB := NewBattle()
	for i, b := range battle.Psychics {
		NB.Psychics[i].Hypothesis = 10 + rand.Intn(99-10+1)
		NB.Psychics[i].Reliability = b.Reliability
	}
	return NB
}

// Calculating the reliability of psychics
func CalculationReliability(battle Battle, number int) *Battle {
	NB := NewBattle()
	for i, b := range battle.Psychics {
		if b.Hypothesis == number {
			NB.Psychics[i].Reliability = 1 + b.Reliability
		} else {
			NB.Psychics[i].Reliability = b.Reliability - 1
		}
		NB.Psychics[i].Hypothesis = b.Hypothesis
	}
	return NB
}
