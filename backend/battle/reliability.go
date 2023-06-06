package battle

// Calculating the reliability of psychics
func CalculationReliability(battle Battle) *Battle {
	NB := NewBattle()
	NB.UserNumber = battle.UserNumber
	for i, b := range battle.Psychics {
		if b.Hypothesis[len(b.Hypothesis)-1] == battle.UserNumber[len(battle.UserNumber)-1] {
			NB.Psychics[i].Reliability = 1 + b.Reliability
		} else {
			NB.Psychics[i].Reliability = b.Reliability - 1
		}
		NB.Psychics[i].Hypothesis = b.Hypothesis
	}
	return NB
}
