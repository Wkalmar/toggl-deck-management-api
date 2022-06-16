package domain

import "errors"

type Card struct {
	value Rank
	suite Shape
}

var shapeToLetterMap map[Shape]string
var letterToShapeMap map[string]Shape
var rankToLetterMap map[Rank]string
var letterToRankMap map[string]Rank

func CreateCard(value Rank, suite Shape) Card {
	return Card{
		value: value,
		suite: suite,
	}
}

func GetCardStringCode(card Card) string {
	return rankToLetterMap[card.value] + shapeToLetterMap[card.suite]
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
