package convertor

import (
	"battle_of_psychics/backend/battle"
)

// The function converts all the guesses of psychics and the numbers that the user entered into the map, so that it is convenient to display it in html
func ConvertvExtrasenceResultToViewStruct(battle battle.Battle) map[int][]int {
	hi := make(map[int][]int)
	for i := 0; i < len(battle.Psychics[0].Hypothesis); i++ {
		sp := make([]int, 0)
		for j := 0; j < 5; j++ {
			sp = append(sp, battle.Psychics[j].Hypothesis[i])
		}
		sp = append(sp, battle.UserNumber[i])
		hi[i] = sp
	}
	return hi
}

// The function converts the authenticity of each psychic into a slice
func ConvertExtrasenceReliabilityToSlice(battle battle.Battle) []int {
	sp := make([]int, 0)
	for _, b := range battle.Psychics {
		sp = append(sp, b.Reliability)
	}
	return sp
}
