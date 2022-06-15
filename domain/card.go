package domain

type Card struct {
	value Rank
	suite Shape
}

var shapeToLetterMap map[Shape]string
var rankToLetterMap map[Rank]string

func CreateCard(value Rank, suite Shape) Card {
	return Card{
		value: value,
		suite: suite,
	}
}

func GetCardStringCode(card Card) string {
	return rankToLetterMap[card.value] + shapeToLetterMap[card.suite]
}
