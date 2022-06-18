package domain

func init() {
	unshuffledCards = generateUnshuffledCards()
	shapeToStringRepresentationMap = make(map[Shape]StringRepresentation)
	shapeToStringRepresentationMap[Spades] = StringRepresentation{code: "S", fullname: "SPADES"}
	shapeToStringRepresentationMap[Clubs] = StringRepresentation{code: "C", fullname: "CLUBS"}
	shapeToStringRepresentationMap[Diamonds] = StringRepresentation{code: "D", fullname: "DIAMONDS"}
	shapeToStringRepresentationMap[Hearts] = StringRepresentation{code: "H", fullname: "HEARTS"}
	rankToStringRepresentationMap = make(map[Rank]StringRepresentation)
	rankToStringRepresentationMap[Ace] = StringRepresentation{code: "A", fullname: "ACE"}
	rankToStringRepresentationMap[Two] = StringRepresentation{code: "2", fullname: "2"}
	rankToStringRepresentationMap[Three] = StringRepresentation{code: "3", fullname: "3"}
	rankToStringRepresentationMap[Four] = StringRepresentation{code: "4", fullname: "4"}
	rankToStringRepresentationMap[Five] = StringRepresentation{code: "5", fullname: "5"}
	rankToStringRepresentationMap[Six] = StringRepresentation{code: "6", fullname: "6"}
	rankToStringRepresentationMap[Seven] = StringRepresentation{code: "7", fullname: "7"}
	rankToStringRepresentationMap[Eight] = StringRepresentation{code: "8", fullname: "8"}
	rankToStringRepresentationMap[Nine] = StringRepresentation{code: "9", fullname: "9"}
	rankToStringRepresentationMap[Ten] = StringRepresentation{code: "10", fullname: "10"}
	rankToStringRepresentationMap[Jack] = StringRepresentation{code: "J", fullname: "JACK"}
	rankToStringRepresentationMap[Queen] = StringRepresentation{code: "Q", fullname: "QUEEN"}
	rankToStringRepresentationMap[King] = StringRepresentation{code: "K", fullname: "KING"}
	letterToShapeMap = make(map[string]Shape)
	letterToShapeMap["S"] = Spades
	letterToShapeMap["C"] = Clubs
	letterToShapeMap["D"] = Diamonds
	letterToShapeMap["H"] = Hearts
	letterToRankMap = make(map[string]Rank)
	letterToRankMap["A"] = Ace
	letterToRankMap["2"] = Two
	letterToRankMap["3"] = Three
	letterToRankMap["4"] = Four
	letterToRankMap["5"] = Five
	letterToRankMap["6"] = Six
	letterToRankMap["7"] = Seven
	letterToRankMap["8"] = Eight
	letterToRankMap["9"] = Nine
	letterToRankMap["10"] = Ten
	letterToRankMap["J"] = Jack
	letterToRankMap["Q"] = Queen
	letterToRankMap["K"] = King
}
