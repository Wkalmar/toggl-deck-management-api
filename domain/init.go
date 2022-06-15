package domain

func init() {
	unshuffledCards = generateUnshuffledCards()
	shapeToLetterMap = make(map[Shape]string)
	shapeToLetterMap[Spades] = "S"
	shapeToLetterMap[Clubs] = "C"
	shapeToLetterMap[Diamonds] = "D"
	shapeToLetterMap[Hearts] = "H"
	rankToLetterMap = make(map[Rank]string)
	rankToLetterMap[Ace] = "A"
	rankToLetterMap[Two] = "2"
	rankToLetterMap[Three] = "3"
	rankToLetterMap[Four] = "4"
	rankToLetterMap[Five] = "5"
	rankToLetterMap[Six] = "6"
	rankToLetterMap[Seven] = "7"
	rankToLetterMap[Eight] = "8"
	rankToLetterMap[Nine] = "9"
	rankToLetterMap[Ten] = "10"
	rankToLetterMap[Jack] = "J"
	rankToLetterMap[Queen] = "Q"
	rankToLetterMap[King] = "K"
}
