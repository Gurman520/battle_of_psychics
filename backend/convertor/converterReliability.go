package convertor

import (
	"battle_of_psychics/backend/battle"
)

func ConvertReliability(battle battle.Battle) []int {
	sp := make([]int, 0)
	for _, b := range battle.Psychics {
		sp = append(sp, b.Reliability)
	}
	return sp
}
