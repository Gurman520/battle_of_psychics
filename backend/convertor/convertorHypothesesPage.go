package convertor

import (
	"battle_of_psychics/backend/battle"
)

func ConvertHypotheses(battle battle.Battle) []int {
	sp := make([]int, 0)
	for _, b := range battle.Psychics {
		sp = append(sp, b.Hypothesis[len(b.Hypothesis)-1])
	}
	return sp
}
