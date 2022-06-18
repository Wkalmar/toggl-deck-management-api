package api

import (
	"toggl-deck-management-api/domain"

	"github.com/google/uuid"
)

type ClosedDeckDTO struct {
	DeckId    uuid.UUID `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining uint8     `json:"remaining"`
}

func CreateClosedDeckDTO(deck domain.Deck) ClosedDeckDTO {
	return ClosedDeckDTO{
		DeckId:    deck.DeckId,
		Shuffled:  deck.Shuffled,
		Remaining: domain.CountRemainingCards(deck),
	}
}
