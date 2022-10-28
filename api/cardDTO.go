package api

import (
	"errors"
	"toggl-deck-management-api/domain"
)

type CardDTO struct {
	Rank  string `json:"value"`
	Shape string `json:"suit"`
	Code  string `json:"code"`
}

type StringRepresentation struct {
	code     string
	fullname string
}

func CreateCardDTO(card domain.Card) CardDTO {
	return CardDTO{
		Rank:  GetRankFullname(card.Rank),
		Shape: GetShapeFullname(card.Shape),
		Code:  GetCardStringCode(card),
	}
}

var shapeToStringRepresentationMap map[domain.Shape]StringRepresentation
var letterToShapeMap map[string]domain.Shape
var rankToStringRepresentationMap map[domain.Rank]StringRepresentation
var letterToRankMap map[string]domain.Rank

func GetCardStringCode(card domain.Card) string {
	return rankToStringRepresentationMap[card.Rank].code + shapeToStringRepresentationMap[card.Shape].code
}

func ParseCardStringCode(code string) (domain.Card, error) {
	if len(code) < 2 || len(code) > 3 {
		return domain.Card{}, errors.New("ParseCardStringCode. Invalid card code")
	}
	rankCode := code[:len(code)-1]
	shapeCode := code[len(code)-1:]
	rank, rankFound := letterToRankMap[rankCode]
	shape, shapeFound := letterToShapeMap[shapeCode]
	if !rankFound || !shapeFound {
		return domain.Card{}, errors.New("ParseCardStringCode. Invalid card code")
	}
	return domain.CreateCard(rank, shape), nil
}

func GetShapeFullname(shape domain.Shape) string {
	return shapeToStringRepresentationMap[shape].fullname
}

func GetRankFullname(rank domain.Rank) string {
	return rankToStringRepresentationMap[rank].fullname
}
