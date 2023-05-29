package api

import (
	"battle_of_psychics/internal/app"
	"battle_of_psychics/openapi/models"
)

func ConceiveToRest(battle app.Battle) []*models.ResponseСonceive {
	m := make([]*models.ResponseСonceive, 5)

	for i, b := range battle.Psychics {
		m[i] = oneHypoToRest(b)
	}

	return m
}

func oneHypoToRest(p app.Psychic) *models.ResponseСonceive {
	return &models.ResponseСonceive{
		Hypothesis: int64(p.Hypothesis),
	}
}

func ResultToRest(battle app.Battle) []*models.ResponseResult {
	r := make([]*models.ResponseResult, 5)

	for i, b := range battle.Psychics {
		r[i] = oneReliabilityToRest(b)
	}

	return r
}

func oneReliabilityToRest(p app.Psychic) *models.ResponseResult {
	n := int64(p.Reliability.Percent)
	return &models.ResponseResult{
		Reliability: &n,
	}
}
