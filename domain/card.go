package domain

import "errors"

type Card struct {
	Value Rank
	Suite Shape
}

var shapeToLetterMap map[Shape]string
var letterToShapeMap map[string]Shape
var rankToLetterMap map[Rank]string
var letterToRankMap map[string]Rank

func CreateCard(value Rank, suite Shape) Card {
	return Card{
		Value: value,
		Suite: suite,
	}
}

func GetCardStringCode(card Card) string {
	return rankToLetterMap[card.Value] + shapeToLetterMap[card.Suite]
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
