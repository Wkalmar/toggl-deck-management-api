package domain

import "errors"

type Card struct {
	Rank  Rank
	Shape Shape
}

type StringRepresentation struct {
	code     string
	fullname string
}

var shapeToStringRepresentationMap map[Shape]StringRepresentation
var letterToShapeMap map[string]Shape
var rankToStringRepresentationMap map[Rank]StringRepresentation
var letterToRankMap map[string]Rank

func CreateCard(rank Rank, shape Shape) Card {
	return Card{
		Rank:  rank,
		Shape: shape,
	}
}

func GetCardStringCode(card Card) string {
	return rankToStringRepresentationMap[card.Rank].code + shapeToStringRepresentationMap[card.Shape].code
}

func ParseCardStringCode(code string) (Card, error) {
	if len(code) < 2 || len(code) > 3 {
		return Card{}, errors.New("ParseCardStringCode. Invalid card code")
	}
	rankCode := code[:len(code)-1]
	shapeCode := code[len(code)-1:]
	rank, rankFound := letterToRankMap[rankCode]
	shape, shapeFound := letterToShapeMap[shapeCode]
	if !rankFound || !shapeFound {
		return Card{}, errors.New("ParseCardStringCode. Invalid card code")
	}
	return CreateCard(rank, shape), nil
}

func GetShapeFullname(shape Shape) string {
	return shapeToStringRepresentationMap[shape].fullname
}

func GetRankFullname(rank Rank) string {
	return rankToStringRepresentationMap[rank].fullname
}
