package convertor

import (
	"battle_of_psychics/backend/battle"
)

func ConvertRank(battle battle.Battle) map[int][]int {
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
