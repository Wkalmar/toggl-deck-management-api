package domain

import "strconv"

func init() {
	unshuffledCards = generateUnshuffledCards()
	shapeToStringRepresentationMap = make(map[Shape]StringRepresentation)
	shapeToStringRepresentationMap[Spades] = StringRepresentation{code: "S", fullname: "SPADES"}
	shapeToStringRepresentationMap[Clubs] = StringRepresentation{code: "C", fullname: "CLUBS"}
	shapeToStringRepresentationMap[Diamonds] = StringRepresentation{code: "D", fullname: "DIAMONDS"}
	shapeToStringRepresentationMap[Hearts] = StringRepresentation{code: "H", fullname: "HEARTS"}
	rankToStringRepresentationMap = make(map[Rank]StringRepresentation)
	rankToStringRepresentationMap[Ace] = StringRepresentation{code: "A", fullname: "ACE"}
	for i := 1; i <= 9; i++ {
		rankToStringRepresentationMap[Rank(i)] = StringRepresentation{code: strconv.Itoa(i + 1), fullname: strconv.Itoa(i + 1)}
	}
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
	for i := 1; i <= 9; i++ {
		letterToRankMap[strconv.Itoa(i+1)] = Rank(i)
	}
	letterToRankMap["J"] = Jack
	letterToRankMap["Q"] = Queen
	letterToRankMap["K"] = King
}
