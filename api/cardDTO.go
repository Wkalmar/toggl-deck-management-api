package api

import "toggl-deck-management-api/domain"

type CardDTO struct {
	Rank  string `json:"value"`
	Shape string `json:"suit"`
	Code  string `json:"code"`
}

func CreateCardDTO(card domain.Card) CardDTO {
	return CardDTO{
		Rank:  domain.GetRankFullname(card.Rank),
		Shape: domain.GetShapeFullname(card.Shape),
		Code:  domain.GetCardStringCode(card),
	}
}
