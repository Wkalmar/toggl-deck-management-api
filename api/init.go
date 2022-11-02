package api

import (
	"strconv"
	"toggl-deck-management-api/domain"
)

func init() {
	shapeToStringRepresentationMap = make(map[domain.Shape]StringRepresentation)
	shapeToStringRepresentationMap[domain.Spades] = StringRepresentation{code: "S", fullname: "SPADES"}
	shapeToStringRepresentationMap[domain.Clubs] = StringRepresentation{code: "C", fullname: "CLUBS"}
	shapeToStringRepresentationMap[domain.Diamonds] = StringRepresentation{code: "D", fullname: "DIAMONDS"}
	shapeToStringRepresentationMap[domain.Hearts] = StringRepresentation{code: "H", fullname: "HEARTS"}
	rankToStringRepresentationMap = make(map[domain.Rank]StringRepresentation)
	rankToStringRepresentationMap[domain.Ace] = StringRepresentation{code: "A", fullname: "ACE"}
	for i := 1; i <= 9; i++ {
		rankToStringRepresentationMap[domain.Rank(i)] = StringRepresentation{code: strconv.Itoa(i + 1), fullname: strconv.Itoa(i + 1)}
	}
	rankToStringRepresentationMap[domain.Jack] = StringRepresentation{code: "J", fullname: "JACK"}
	rankToStringRepresentationMap[domain.Queen] = StringRepresentation{code: "Q", fullname: "QUEEN"}
	rankToStringRepresentationMap[domain.King] = StringRepresentation{code: "K", fullname: "KING"}
	letterToShapeMap = make(map[string]domain.Shape)
	letterToShapeMap["S"] = domain.Spades
	letterToShapeMap["C"] = domain.Clubs
	letterToShapeMap["D"] = domain.Diamonds
	letterToShapeMap["H"] = domain.Hearts
	letterToRankMap = make(map[string]domain.Rank)
	letterToRankMap["A"] = domain.Ace
	for i := 1; i <= 9; i++ {
		letterToRankMap[strconv.Itoa(i+1)] = domain.Rank(i)
	}
	letterToRankMap["J"] = domain.Jack
	letterToRankMap["Q"] = domain.Queen
	letterToRankMap["K"] = domain.King
}
